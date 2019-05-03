import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BlogsComponent } from './blogs/blogs.component';
import { ArticlesComponent } from './articles/articles.component';

const routes: Routes = [
  { path: '', component: BlogsComponent},
  { path: 'blog/:id/articles', component: ArticlesComponent},
  {path: '**', redirectTo: 'blogs'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
