import { Component, AfterViewInit, ElementRef } from '@angular/core';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { Router } from '@angular/router';

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
  textGenerationComplete = false;
  backgroundSquareVisible = false; 

  constructor(private elementRef: ElementRef,  private router: Router) {}

  ngAfterViewInit() {
    setTimeout(() => {
      this.generateText();
    }, 300); // slight delay to allow the page to load before the text starts generating
  }

  // Function to generate the text character by character
  generateText() {
    const fullText = "Welcome to MealDealz! Please create an account or login.";
    let currentIndex = 0;

    const interval = setInterval(() => {
      if (currentIndex >= fullText.length) {
        clearInterval(interval);
        this.textGenerationComplete = true;
        setTimeout(() => {
          this.backgroundSquareVisible = true;
        }, 500);
        return;
      }

      this.animatedText += fullText[currentIndex];
      currentIndex++;

      // open and close mouth
      if (currentIndex % 2 === 0)
        this.toggleAnimation();

    }, 50); // You can adjust the speed of text generation by changing the interval time (in milliseconds).

    // resets the mouth to the open position
    this.isMouthTalking = currentIndex % 2 === 1;
  }

  // Function to toggle the animation state
  toggleAnimation() {
    this.isMouthTalking = !this.isMouthTalking;
  }

  navigateToLoginPage() {
    this.router.navigate(['/Login']);
  }

  navigateToSignupPage() {
    this.router.navigate(['/Signup']);
  }

}