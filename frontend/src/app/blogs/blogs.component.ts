import { Component, OnInit } from '@angular/core';
import { WebService } from '../services/web.service';

@Component({
  selector: 'app-blogs',
  templateUrl: './blogs.component.html',
  styleUrls: ['./blogs.component.css']
})
export class BlogsComponent implements OnInit {
  public blogs: [] = [];

  constructor(private webService: WebService) { }

  ngOnInit() {
    this.loadBlogs();
  }

  loadBlogs() {
    this.webService.getBlogs().subscribe((data: []) => {
      this.blogs = data;
    });
  }

}
