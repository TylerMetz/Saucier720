import { Component } from '@angular/core';
import { trigger, state, style, transition, animate } from '@angular/animations';

@Component({
    selector: 'app-animated-logo',
    templateUrl: './animated-logo.component.html',
    styleUrls: ['./animated-logo.component.scss'],
    animations: [
      trigger('mouthTalking', [
        state('false', style({ transform: 'scaleY(1)' })),
        state('true', style({ transform: 'scaleY(0.7)' })),
        transition('false <=> true', animate('200ms ease-out')),
      ]),
    ],
  })
  export class AnimatedLogoComponent {
    isMouthTalking = false;
  
    // Function to toggle the animation state
    toggleAnimation() {
      this.isMouthTalking = !this.isMouthTalking;
    }
  }
  