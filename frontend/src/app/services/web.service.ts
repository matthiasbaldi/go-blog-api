import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class WebService {

  constructor(private http: HttpClient) { }

  public getBlogs() {
    return this.http
      .get(`/api/v1/blogs`);
  }

  public getArticles(blogId: string) {
    return this.http
      .get(`/api/v1/blogs/${blogId}/articles`);
  }
}
