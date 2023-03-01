import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PhotoUpComponent } from './photo-upload.component';

const routes: Routes = [
  {
    path: '',
    component: PhotoUpComponent,
    children: [
      {
        path: 'photo-upload',
        component: PhotoUpComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PhotoUpRoutingModule { }
