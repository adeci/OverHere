import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ReturninguserComponent } from './returning-user.component';

const routes: Routes = [
  {
    path: '',
    component: ReturninguserComponent,
    children: [
      {
        path: 'returning-user',
        component: ReturninguserComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ReturninguserRoutingModule { }