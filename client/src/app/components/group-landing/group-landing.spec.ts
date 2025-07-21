import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupLanding } from './group-landing';

describe('GroupLanding', () => {
  let component: GroupLanding;
  let fixture: ComponentFixture<GroupLanding>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GroupLanding]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GroupLanding);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
