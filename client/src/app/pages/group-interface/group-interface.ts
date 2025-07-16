import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { GroupService } from '../../services/group/group';
import { GroupMember } from '../../types/group';
import { AuthService } from '../../services/auth/auth';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { CommonModule } from '@angular/common';

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
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.groupCode = this.route.snapshot.paramMap.get('groupCode') ?? '';
    this.fetchGroupMembers();
  }

  fetchGroupMembers(): void {
    const token = this.authService.getToken();
    if (!token) {
      console.error('No token found. User may not be authenticated.');
      return;
    }

    this.groupService.fetchGroupMembers(token, this.groupCode).subscribe({
      next: (res) => {
        this.members = res.data ?? [];
      },
      error: (err) => {
        console.error('Failed to fetch group members:', err);
      }
    });
  }

  copyGroupCode(): void {
    navigator.clipboard.writeText(this.groupCode).then(() => {
      console.log('Group code copied!');
    }).catch(err => {
      console.error('Clipboard error:', err);
    });
  }

  shareGroupCode(): void {
    if (navigator.share) {
      navigator.share({
        title: 'Join my SwipEats group',
        text: `Use this code to join: ${this.groupCode}`,
        url: window.location.href
      }).catch((err) => console.error('Share failed:', err));
    } else {
      this.copyGroupCode();
    }
  }
}
