import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from "./components/login/login.component";
import {MessagingComponent} from "./components/messaging/messaging.component";

const routes: Routes = [
  {
    path: '',
    component: LoginComponent
  },
  {
    path: 'msg',
    component: MessagingComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ChatRoutingModule { }
