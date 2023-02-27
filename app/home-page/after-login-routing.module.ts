import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AfterloginComponent } from './after-login.component';

const routes: Routes = [
  {
    path: '',
    component: AfterloginComponent,
    children: [
      {
        path: 'home',
        component: AfterloginComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AfterloginRoutingModule { }