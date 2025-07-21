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

@Component({
  selector: 'app-group-interface',
  imports: [CommonModule, NavigationBar, GroupLanding, GroupStarted],
  templateUrl: './group-interface.html',
  styleUrl: './group-interface.css'
})
export class GroupInterface implements OnInit {
  groupLanding: string = 'loading'; // Default to landing view

  groupCode: string = '';
  members: GroupMember[] = [];
  userId: number | null = null;
  isOwner: boolean = false;
  baseImageUrl: string = `${API_URL_V1}/uploads/`;
  imageLoadFailed: { [email: string]: boolean } = {};

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private groupService: GroupService,
    private authService: AuthService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.userId = this.authService.getUserId();
    this.groupCode = this.route.snapshot.paramMap.get('groupCode') ?? '';
    if (!this.groupCode) {
      this.toastr.error('Group code is required', 'Error');
      return;
    }
      this.groupService.connectWebSocket(
        this.authService.getToken() || '',
        this.groupCode,
        (data) => {
          if (data.group_status && data.group_status === 'active') {
            this.groupLanding = 'started';
          }
          else if (data.group_status && data.group_status === 'waiting') {
            this.groupLanding = 'landing';
          }
          else {
            this.toastr.error('Invalid group status', 'Error');
            return;
          }

          if (data.type === 'members_update') {
            this.members = [...data.members.map((member: GroupMember) => {
              if (member.user_id === this.userId) {
                this.isOwner = member.is_owner;
              }

              return {
                ...member,
                name: member.name || 'User',
              }
            })];
          }
          else if (data.type === 'group_session_started') {
            // Handle group session start, e.g., navigate to session page or show a message
            this.toastr.success('Group session started successfully', 'Success');
            this.groupLanding = 'started';
          }
          else if (data.type === 'group_session_ended') {
            // Handle group session end, e.g., navigate back or show a message
            this.router.navigate(['/lobby']);
            this.toastr.success('Group session ended', 'Success');
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
}
