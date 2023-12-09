import { Component, OnDestroy, OnInit, signal } from '@angular/core';
import { Subscription } from 'rxjs';
import { IFileTreeDto } from '../../models/fileTreeDto';
import { FetchFileTreeService } from '../../services/fetch-file-tree.service';
import { NavigationTreeComponent } from './navigation-tree/navigation-tree.component';

@Component({
  selector: 'app-navigation',
  standalone: true,
  imports: [NavigationTreeComponent],
  templateUrl: './navigation.component.html',
  styleUrl: './navigation.component.css',
})
export class NavigationComponent implements OnInit, OnDestroy {
  public fileTrees = signal<IFileTreeDto[]>([]);
  public isLoading = signal<boolean>(false);
  public error = signal<any>(undefined);
  private subscription: Subscription | undefined;

  constructor(private fetchFileTreeService: FetchFileTreeService) {}

  ngOnInit() {
    this.getFileTrees();
  }

  ngOnDestroy(): void {
    this.subscription?.unsubscribe();
  }

  getFileTrees(): void {
    this.isLoading.update((_) => true);
    this.subscription = this.fetchFileTreeService.getFileTrees().subscribe({
      complete: () => this.isLoading.update((_) => false),
      next: (fileTree) => this.fileTrees.update((_) => fileTree),
      error: this.error.update,
    });
  }
}
