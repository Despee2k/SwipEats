import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupInterface } from './group-interface';

describe('GroupInterface', () => {
  let component: GroupInterface;
  let fixture: ComponentFixture<GroupInterface>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GroupInterface]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GroupInterface);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
