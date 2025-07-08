import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Lobby } from './lobby';

describe('Lobby', () => {
  let component: Lobby;
  let fixture: ComponentFixture<Lobby>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Lobby]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Lobby);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
