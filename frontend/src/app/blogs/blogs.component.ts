import { Component, OnInit } from '@angular/core';
import { WebService } from '../services/web.service';

@Component({
  selector: 'app-blogs',
  templateUrl: './blogs.component.html',
  styleUrls: ['./blogs.component.css']
})
export class BlogsComponent implements OnInit {
  public blogs: [] = [];
  public newBlog: any;

  constructor(private webService: WebService) {
    this.setDefaults();
  }

  ngOnInit() {
    this.loadBlogs();
  }

  loadBlogs() {
    this.webService.getBlogs().subscribe((data: []) => {
      this.blogs = data;
    });
  }

  public addBlog() {
    this.newBlog.slug = this.newBlog.title.toUpperCase().substring(1, 4);
    console.log(this.newBlog);
    this.webService.createBlog(this.newBlog).subscribe((data: any) => {
      this.setDefaults();
      alert(data);
      this.loadBlogs();
    });
  }

  public removeBlog(id: string) {
    this.webService.removeBlog(id).subscribe((data: any) => {
      alert('successfully removed');
      this.loadBlogs();
    });
  }

  private setDefaults() {
    this.newBlog = {
      title: '',
      description: '',
      author: '',
      slug: ''
    };
  }

}
