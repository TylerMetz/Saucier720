import { Component, AfterViewInit, HostListener, ElementRef } from '@angular/core';
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

  constructor(private elementRef: ElementRef) {}

  ngAfterViewInit() {
    this.generateText();
  }

  // Function to generate the text character by character
  generateText() {
    const fullText = "Welcome to MealDealz! Please create an account or login.";
    let currentIndex = 0;

    const interval = setInterval(() => {
      if (currentIndex >= fullText.length) {
        clearInterval(interval);
        return;
      }

      this.animatedText += fullText[currentIndex];
      currentIndex++;

      // open and close mouth
      if (currentIndex % 2 === 0)
        this.toggleAnimation();

      // check if text wrapping needs enabled
      /*
      if (currentIndex > 30){
        this.elementRef.nativeElement.querySelector('.text-container p').style.whiteSpace = 'normal';
      }
      */

    }, 50); // You can adjust the speed of text generation by changing the interval time (in milliseconds).

     // resets the mouth to the open position
    this.isMouthTalking = currentIndex % 2 === 1;
  }

  // Function to toggle the animation state
  toggleAnimation() {
    this.isMouthTalking = !this.isMouthTalking;
  }

}
