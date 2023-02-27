import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthRoutingModule } from '../auth/auth-routing.module';
import { ReturninguserComponent } from './returning-user.component';
import { ReturninguserRoutingModule } from './returning-user-routing.module';
import { Router, RouterModule } from '@angular/router';


@NgModule({
  declarations: [ReturninguserComponent],
  imports: [
    CommonModule,
    ReturninguserRoutingModule,
    RouterModule
  ]
})
export class ReturninguserModule { }