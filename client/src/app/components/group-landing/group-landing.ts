import { Component, Input } from '@angular/core';
import { GroupMember } from '../../types/group';
import { API_URL_V1 } from '../../utils/constant';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { GroupService } from '../../services/group/group';
import { ToastrService } from 'ngx-toastr';
import { getRoundedDownHour } from '../../utils/general';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-group-landing',
  imports: [CommonModule, RouterLink],
  templateUrl: './group-landing.html',
  styleUrl: './group-landing.css'
})
export class GroupLanding {
  @Input() groupCode!: string;
  @Input() members!: GroupMember[];
  @Input() userId!: number | null;
  @Input() isOwner!: boolean;
  @Input() endGroup!: () => void;
  @Input() leaveGroup!: () => void;
  @Input() startGroup!: () => void;
  baseImageUrl: string = `${API_URL_V1}/uploads/`;
  imageLoadFailed: { [email: string]: boolean } = {};

  constructor(
    private toastr: ToastrService
  ) {}

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

  getEncodedImageUrl(email: string): string {
    return this.baseImageUrl + encodeURIComponent(email) + '?t=' + getRoundedDownHour(); // Prevent caching
  }
}
