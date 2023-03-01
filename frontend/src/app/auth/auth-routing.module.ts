import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPagesComponent } from '../pages/login-pages/login-pages.component';

const routes: Routes = [
  {
    path: '',
    component: LoginPagesComponent,
    children: [
      {
        path: 'auth',
        component: LoginPagesComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
