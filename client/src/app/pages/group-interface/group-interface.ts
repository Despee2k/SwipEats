import { Component, OnInit, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { GroupService } from '../../services/group/group';
import { GroupMember } from '../../types/group';
import { AuthService } from '../../services/auth/auth';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { CommonModule } from '@angular/common';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-group-interface',
  imports: [CommonModule, RouterLink, NavigationBar],
  templateUrl: './group-interface.html',
  styleUrl: './group-interface.css'
})
export class GroupInterface implements OnInit {
  groupCode: string = '';
  members: GroupMember[] = [];

  constructor(
    private route: ActivatedRoute,
    private groupService: GroupService,
    private authService: AuthService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
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
            this.members = [...data.members];
          }
          else if (data.type === 'group_session_started') {
            // Handle group session start, e.g., navigate to session page or show a message
            this.toastr.success('Group session started successfully', 'Success');
          }
          else if (data.type === 'group_session_ended') {
            // Handle group session end, e.g., navigate back or show a message
            this.toastr.info('Group session ended', 'Info');
          }
        },
        (err) => {
          console.error('WebSocket error:', err);
          this.toastr.error('Failed to connect to group session', 'Error');
        }
      );
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
      }).catch((err) => this.toastr.error('Share failed', 'Error'));
    } else {
      this.copyGroupCode();
    }
  }
}
