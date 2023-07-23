import { Component } from '@angular/core';
import { trigger, state, style, transition, animate } from '@angular/animations';

@Component({
    selector: 'app-landing-page',
    templateUrl: './landing-page.component.html',
    styleUrls: ['./landing-page.component.scss'],
    animations: [
      trigger('mouthTalking', [
        state('false', style({ transform: 'scaleY(1)' })),
        state('true', style({ transform: 'scaleY(0.7)' })),
        transition('false <=> true', animate('200ms ease-out')),
      ]),
    ],
  })
  export class LandingPageComponent {
    isMouthTalking = false;
  
    // Function to toggle the animation state
    toggleAnimation() {
      this.isMouthTalking = !this.isMouthTalking;
    }
  }
  