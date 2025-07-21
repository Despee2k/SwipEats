import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Swiping } from './swiping';

describe('Swiping', () => {
  let component: Swiping;
  let fixture: ComponentFixture<Swiping>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Swiping]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Swiping);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
