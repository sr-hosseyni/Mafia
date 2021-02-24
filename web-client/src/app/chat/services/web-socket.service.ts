import { Injectable } from '@angular/core';

/**
 * Best Practice with Observabales
 * @see https://tutorialedge.net/typescript/angular/angular-websockets-tutorial/
 *
 * websocket tutorial
 * @see https://www.npmjs.com/package/ng-ws
 */

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {
  ws: WebSocket;
  constructor() { }

  public connect(url, listener)  {
    this.ws = new WebSocket(url);
    this.ws.addEventListener('message', listener);
  }

  send(data: string) {
    this.ws.send(data);
  }
}
