import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PhotoLibComponent } from './photo-lib.component';

const routes: Routes = [
  {
    path: '',
    component: PhotoLibComponent,
    children: [
      {
        path: 'photo-library',
        component: PhotoLibComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PhotoLibRoutingModule { }