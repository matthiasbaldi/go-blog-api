import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <nav class="navbar navbar-dark bg-dark">
      <span class="navbar-brand mb-0 h1 pointer" routerLink="/">Blog Management</span>
    </nav>
    <div class="container">
      <router-outlet></router-outlet>
    </div>
  `,
  styles: []
})
export class AppComponent {
  title = 'blog-client';
}
