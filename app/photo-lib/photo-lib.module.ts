import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PhotoLibRoutingModule } from './photo-lib-routing.module';
import { PhotoLibComponent } from './photo-lib.component';
import { Router, RouterModule } from '@angular/router';


@NgModule({
  declarations: [PhotoLibComponent],
  imports: [
    CommonModule,
    PhotoLibRoutingModule,
    RouterModule
  ]
})
export class PhotoLibModule { }