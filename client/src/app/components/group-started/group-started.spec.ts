import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupStarted } from './group-started';

describe('GroupStarted', () => {
  let component: GroupStarted;
  let fixture: ComponentFixture<GroupStarted>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GroupStarted]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GroupStarted);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
