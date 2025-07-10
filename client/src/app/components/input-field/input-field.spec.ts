import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InputField } from './input-field';

describe('InputField', () => {
  let component: InputField;
  let fixture: ComponentFixture<InputField>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InputField]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InputField);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
