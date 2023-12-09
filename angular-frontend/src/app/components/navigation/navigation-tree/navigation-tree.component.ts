import { Component, Input } from '@angular/core';
import { NavigationLeafComponent } from '../navigation-leaf/navigation-leaf.component';
import { IFileTreeDto } from '../../../models/fileTreeDto';

@Component({
  selector: 'app-navigation-tree',
  standalone: true,
  imports: [NavigationLeafComponent],
  templateUrl: './navigation-tree.component.html',
  styleUrl: './navigation-tree.component.css',
})
export class NavigationTreeComponent {
  @Input({ required: true }) fileTree!: IFileTreeDto;
}
