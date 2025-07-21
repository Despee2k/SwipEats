import { Component, OnInit, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { GroupService } from '../../services/group/group';
import { GroupMember } from '../../types/group';
import { AuthService } from '../../services/auth/auth';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { CommonModule } from '@angular/common';
import { ToastrService } from 'ngx-toastr';
import { API_URL_V1 } from '../../utils/constant';

@Component({
  selector: 'app-group-interface',
  imports: [CommonModule, RouterLink, NavigationBar],
  templateUrl: './group-interface.html',
  styleUrl: './group-interface.css'
})
export class GroupInterface implements OnInit {
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

  handleImageError(email: string): void {
    this.imageLoadFailed[email] = true;
  }

  copyGroupCode(): void {
    navigator.clipboard.writeText(this.groupCode).then(() => {
      this.toastr.success('Group code copied!', 'Success');
    }).catch(err => {
      this.toastr.error('Failed to copy group code', 'Error');
    });
  }

  shareGroupCode(): void {
    if (navigator.share) {
      navigator.share({
        title: 'Join my SwipEats group',
        text: `Use this code to join: ${this.groupCode}`,
        url: window.location.href
      }).catch((err) => console.error('Share failed'));
    } else {
      this.copyGroupCode();
    }
  }

  endGroup(): void {
    this.router.navigate(['/lobby']);
    this.groupService.endGroup();
  }

  leaveGroup(): void {
    this.router.navigate(['/lobby']);
    this.toastr.success('You have left the group', 'Success');
    this.groupService.leaveGroup();
  }
}
