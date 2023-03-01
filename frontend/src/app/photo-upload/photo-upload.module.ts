import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PhotoUpRoutingModule } from './photo-upload-routing.module';
import { PhotoUpComponent } from './photo-upload.component';


@NgModule({
  declarations: [PhotoUpComponent],
  imports: [
    CommonModule,
    PhotoUpRoutingModule
  ]
})
export class PhotoUpModule { }