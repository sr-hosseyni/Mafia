import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ChatRoutingModule } from './chat-routing.module';
import { LoginComponent } from './components/login/login.component';
import { MessagingComponent } from './components/messaging/messaging.component';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {HttpClientModule} from "@angular/common/http";

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule,
    ChatRoutingModule
  ],
  declarations: [LoginComponent, MessagingComponent],
  // providers: [HttpClientModule]
})
export class ChatModule { }
