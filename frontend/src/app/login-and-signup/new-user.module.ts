import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthRoutingModule } from '../auth/auth-routing.module';
import { NewuserComponent } from './new-user.component';
import { NewuserRoutingModule } from './new-user-routing.module';
import { Router, RouterModule } from '@angular/router';


@NgModule({
  declarations: [NewuserComponent],
  imports: [
    CommonModule,
    NewuserRoutingModule,
    RouterModule
  ]
})
export class NewuserModule { }