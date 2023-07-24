import { Component, AfterViewInit } from '@angular/core';
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
export class AnimatedLogoComponent implements AfterViewInit {
  isMouthTalking = false;
  animatedText = "";

  ngAfterViewInit() {
    this.generateText();
  }

  // Function to generate the text character by character
  generateText() {
    const fullText = "Welcome to MealDealz! Please create an account or login";
    let currentIndex = 0;

    const interval = setInterval(() => {
      if (currentIndex >= fullText.length) {
        clearInterval(interval);
        return;
      }

      this.animatedText += fullText[currentIndex];
      currentIndex++;
      if (currentIndex % 2 === 0)
        this.toggleAnimation();
    }, 50); // You can adjust the speed of text generation by changing the interval time (in milliseconds).
    this.isMouthTalking = true; // resets the mouth to the open position
  }

  // Function to toggle the animation state
  toggleAnimation() {
    this.isMouthTalking = !this.isMouthTalking;
  }
}
