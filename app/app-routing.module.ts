import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule, Routes } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';

import { LoginPagesComponent } from './pages/login-pages/login-pages.component';
import { MapComponent } from './map/map.component';
import { NewuserPagesComponent } from './pages/new-user-pages/new-user-pages.component';
import { ReturninguserPagesComponent } from './pages/returning-user-pages/returning-user-pages.component';
import { AfterloginPagesComponent } from './pages/after-login-pages/after-login-pages.component';
import { PhotoUpPagesComponent } from './pages/photo-upload-pages/photo-upload-pages.component';
import { PhotoLibPagesComponent } from './pages/photo-lib-pages/photo-lib-pages.component';

const routes: Routes = [
  {
    path: '',
    component: LoginPagesComponent,
    children: [
      {
        path: '',
        redirectTo: '/login',
        pathMatch: 'full'
      },
      {
        path: 'login',
        loadChildren: () => import('./login/login.module').then(m => m.LoginModule)
      }
    ]
  },

  {
    path: '',
    component: MapComponent,
    children: [
      {
        path: '',
        redirectTo: '/map',
        pathMatch: 'full'
      },
      {
        path: 'map',
        loadChildren: () => import('./map/map.module').then(m => m.MapModule)
      }
    ]
  },

  {
    path:'',
    component: NewuserPagesComponent,
    children: [
      {
        path: '',
        redirectTo: '/new-user',
        pathMatch:"full"
      },
      {
        path:'new-user',
        loadChildren: () => import('./login-and-signup/new-user.module').then(m => m.NewuserModule)
      }
    ]
  },

  {
    path:'',
    component: ReturninguserPagesComponent,
    children: [
      {
        path: '',
        redirectTo: '/returning-user',
        pathMatch:"full"
      },
      {
        path: 'returning-user',
        loadChildren: () => import('./login-and-signup/returning-user.module').then(m => m.ReturninguserModule)
      }
    ]
  },

  {
    path:'',
    component: AfterloginPagesComponent,
    children: [
      {
        path: '',
        redirectTo: '/home',
        pathMatch:"full"
      },
      {
        path: 'home',
        loadChildren: () => import('./home-page/after-login.module').then(m => m.AfterloginModule)
      }
    ]
  },

  {
    path:'',
    component: PhotoUpPagesComponent,
    children: [
      {
        path: '',
        redirectTo: '/photo-upload',
        pathMatch:"full"
      },
      {
        path: 'photo-upload',
        loadChildren: () => import('./photo-upload/photo-upload.module').then(m => m.PhotoUpModule)
      }
    ]
  },

  {
    path:'',
    component: PhotoLibPagesComponent,
    children: [
      {
        path: '',
        redirectTo: '/photo-library',
        pathMatch:"full"
      },
      {
        path: 'photo-library',
        loadChildren: () => import('./photo-lib/photo-lib.module').then(m => m.PhotoLibModule)
      }
    ]
  },
]

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forRoot(routes),
    RouterTestingModule
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
