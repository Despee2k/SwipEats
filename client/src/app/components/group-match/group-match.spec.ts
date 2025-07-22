import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupMatch } from './group-match';

describe('GroupMatch', () => {
  let component: GroupMatch;
  let fixture: ComponentFixture<GroupMatch>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GroupMatch]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GroupMatch);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
