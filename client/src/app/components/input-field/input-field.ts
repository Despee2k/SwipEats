import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-input-field',
  imports: [],
  templateUrl: './input-field.html',
  styleUrl: './input-field.css'
})
export class InputField {
  @Input() type: string = 'text';
  @Input() id: string = '';
  @Input() name: string = '';
  @Input() required: boolean = false;
  @Input() placeholder: string = '';
  @Input() modelValue: string = '';
  @Input() label: string = '';

  @Output() modelValueChange = new EventEmitter<string>();

  onInput(event: Event) {
    const value = (event.target as HTMLInputElement).value;
    this.modelValueChange.emit(value);
  }
}
