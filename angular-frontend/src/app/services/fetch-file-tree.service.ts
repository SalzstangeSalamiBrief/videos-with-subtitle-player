import { Injectable } from '@angular/core';
import { IFileTreeDto } from '../models/fileTreeDto';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class FetchFileTreeService {
  // TODO remove proxy file proxy.conf.json and remove call of that file from ng serve command and GET URL VIA ENV_FILE
  private path = '/api/file-tree';

  constructor(private httpClient: HttpClient) {}

  getFileTrees(): Observable<IFileTreeDto[]> {
    return this.httpClient.get<IFileTreeDto[]>(this.path);
  }
}
