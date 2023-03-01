import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AfterloginRoutingModule } from './after-login-routing.module';
import { AfterloginComponent } from './after-login.component';
import { RouterModule } from '@angular/router';


@NgModule({
  declarations: [AfterloginComponent],
  imports: [
    CommonModule,
    AfterloginRoutingModule,
    RouterModule,
  ]
})
export class AfterloginModule { }