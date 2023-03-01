import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NewuserComponent } from './new-user.component';

const routes: Routes = [
  {
    path: '',
    component: NewuserComponent,
    children: [
      {
        path: 'new-user',
        component: NewuserComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class NewuserRoutingModule { }