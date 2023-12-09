import { Component, Input } from '@angular/core';
import { IAudioFileDto } from '../../../models/audioFileDto';

@Component({
  selector: 'app-navigation-leaf',
  standalone: true,
  imports: [],
  templateUrl: './navigation-leaf.component.html',
  styleUrl: './navigation-leaf.component.css',
})
export class NavigationLeafComponent {
  @Input({ required: true }) audioFile!: IAudioFileDto;
}
