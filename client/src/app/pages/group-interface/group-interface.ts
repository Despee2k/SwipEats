import { Component, OnInit, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { GroupService } from '../../services/group/group';
import { GroupMember } from '../../types/group';
import { AuthService } from '../../services/auth/auth';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { CommonModule } from '@angular/common';
import { ToastrService } from 'ngx-toastr';
import { API_URL_V1 } from '../../utils/constant';
import { getRoundedDownHour } from '../../utils/general';
import { GroupLanding } from '../../components/group-landing/group-landing';
import { GroupStarted } from '../../components/group-started/group-started';
import { GroupRestaurant } from '../../types/restaurants';
import { GroupMatch } from '../../components/group-match/group-match';

@Component({
  selector: 'app-group-interface',
  imports: [CommonModule, NavigationBar, GroupLanding, GroupStarted, GroupMatch],
  templateUrl: './group-interface.html',
  styleUrl: './group-interface.css'
})
export class GroupInterface implements OnInit {
  groupLanding: string = 'loading'; // Default to landing view

  groupCode: string = '';
  members: GroupMember[] = [];
  userId: number | null = null;
  token: string | null = null;
  isOwner: boolean = false;
  baseImageUrl: string = `${API_URL_V1}/uploads/`;
  imageLoadFailed: { [email: string]: boolean } = {};
  alreadyVoted: boolean = false;
  finalRestaurant: GroupRestaurant | null = null;

  groupRestaurants: GroupRestaurant[] = [];
  votes: { [restaurantId: number]: boolean } = {};

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private groupService: GroupService,
    private authService: AuthService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.userId = this.authService.getUserId();
    this.token = this.authService.getToken();
    this.groupCode = this.route.snapshot.paramMap.get('groupCode') ?? '';
    if (!this.groupCode) {
      this.toastr.error('Group code is required', 'Error');
      return;
    }

    if (!this.token) {
      return;
    }

      this.groupService.connectWebSocket(
        this.token,
        this.groupCode,
        (data) => {
          // Update group page based on group status
          if (data.group_status && data.group_status === 'active') {
            this.groupLanding = 'started';
          }
          else if (data.group_status && data.group_status === 'waiting') {
            this.groupLanding = 'landing';
          }

          // Update group members and restaurants
          if (data.is_finished_swiping !== undefined) {
            this.alreadyVoted = data.is_finished_swiping;
          }

          if (data.type === 'members_update') {
            console.log('Members update received:', data);
            this.members = [...data.members.map((member: GroupMember) => {
              if (member.user_id === this.userId) {
                this.isOwner = member.is_owner;
              }

              return {
                ...member,
                name: member.name || 'User',
              }
            })];

            if (data.group_restaurants) {
              this.groupRestaurants = data.group_restaurants.filter((restaurant: GroupRestaurant) => {
                return this.votes[restaurant.id] === undefined;
              });
            }
            else {
              console.log('No restaurants found for this group');
            }
          }
          else if (data.type === 'group_session_started') {
            // Handle group session start, e.g., navigate to session page or show a message
            this.toastr.success('Group session started successfully', 'Success');
            this.groupLanding = 'started';

            if (data.group_restaurants) {
              this.groupRestaurants = data.group_restaurants.filter((restaurant: GroupRestaurant) => {
                return this.votes[restaurant.id] === undefined;
              });
            }
            else {
              console.log('No restaurants found for this group');
            }
          }
          else if (data.type === 'group_session_ended') {
            // Handle group session end, e.g., navigate back or show a message
            if (data.most_liked_group_restaurant) {
              this.finalRestaurant = data.most_liked_group_restaurant;
              this.groupLanding = 'match';
            }
            else {
              this.toastr.info('Group session ended without a match', 'Info');
              this.router.navigate(['/lobby']);
            }
          }
        },
        (err) => {
          this.router.navigate(['/lobby']);
          console.error('WebSocket error:', err);
        }
      );
  }

  endGroup = (): void => {
    this.router.navigate(['/lobby']);
    this.groupService.endGroup();
  }

  leaveGroup = (): void => {
    this.router.navigate(['/lobby']);
    this.toastr.success('You have left the group', 'Success');
    this.groupService.leaveGroup();
  }

  startGroup = (): void => {
    this.groupService.startGroupSession();
  }

  handleVote = (groupRestaurantId: number, vote: boolean): void => {
    if (this.votes[groupRestaurantId] !== undefined) {
      this.toastr.warning('You have already voted for this restaurant', 'Warning');
      return;
    }

    this.votes[groupRestaurantId] = vote;
    console.log(this.votes);
  }

  handleSubmitVotes = (): void => {
    if (Object.keys(this.votes).length === 0) {
      this.toastr.warning('You must vote for at least one restaurant', 'Warning');
      return;
    }

    this.groupService.submit_swipes(this.votes);
    this.toastr.success('Votes submitted successfully', 'Success');
    this.groupRestaurants = [];
    this.votes = {};
  }
}
