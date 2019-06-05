import { Component, OnInit } from '@angular/core';
import { WebService } from '../services/web.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-articles',
  templateUrl: './articles.component.html',
  styleUrls: ['./articles.component.css']
})
export class ArticlesComponent implements OnInit {
  public articles: [] = [];
  public selectedBlog: string;
  public blog: null;

  constructor(private route: ActivatedRoute, private webService: WebService) {
    this.route.params.subscribe(params => {
      this.selectedBlog = params.id;
    });
  }

  ngOnInit() {
    this.loadArticles(this.selectedBlog);
    this.loadBlog(this.selectedBlog);
  }

  loadBlog(id: string) {
    this.webService.getBlog(id).subscribe((data: any) => {
      this.blog = data;
    });
  }

  loadArticles(id: string) {
    this.webService.getArticles(id).subscribe((data: []) => {
      this.articles = data;
    });
  }

}
