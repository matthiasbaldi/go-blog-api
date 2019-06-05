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

  public getBlog(id: string) {
    return this.http
      .get(`/api/v1/blogs/${id}`);
  }

  public createBlog(blog: any) {
    return this.http
      .post(`/api/v1/blogs`, blog);
  }

  removeBlog(id: string) {
    return this.http
      .delete(`/api/v1/blogs/${id}`);
  }

  public getArticles(blogId: string) {
    return this.http
      .get(`/api/v1/blogs/${blogId}/articles`);
  }
}
