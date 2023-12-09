import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { IFileTreeDto } from '../../models/fileTreeDto';
import { FetchFileTreeService } from '../../services/fetch-file-tree.service';

@Component({
  selector: 'app-navigation',
  standalone: true,
  imports: [],
  templateUrl: './navigation.component.html',
  styleUrl: './navigation.component.css',
})
export class NavigationComponent implements OnInit, OnDestroy {
  public fileTree: IFileTreeDto | undefined = undefined;
  public isLoading = false;
  public error: any;
  private subscription: Subscription | undefined;

  constructor(private fetchFileTreeService: FetchFileTreeService) {}

  ngOnInit() {
    console.log('oninit');
    this.getFileTree();
  }

  ngOnDestroy(): void {
    this.subscription?.unsubscribe();
  }

  getFileTree(): void {
    this.isLoading = true;
    this.subscription = this.fetchFileTreeService.getFileTree().subscribe({
      complete: () => (this.isLoading = false),
      next: (fileTree) => (this.fileTree = fileTree),
      error: (error) => (this.error = error),
    });
  }
}
