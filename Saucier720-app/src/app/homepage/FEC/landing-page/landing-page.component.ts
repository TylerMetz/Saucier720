import { Component, ElementRef, ViewChild, Renderer2 } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';

@Component({
    selector: 'app-landing-page',
    templateUrl: './landing-page.component.html',
    styleUrls: ['./landing-page.component.scss']
  })

  export class LandingPageComponent{
    buttonsVisible: boolean = false;
    isComponentReady = false;


    onLogoGenerationComplete(generationComplete: boolean) {
      if (generationComplete) {
        // The logo generation is complete, you can now trigger the buttons generation or their appearance.
        // For example, you can set a flag to control their visibility:
        this.buttonsVisible = true;
        this.isComponentReady = true;
      }
    }

    clickedButton: string | null = null;

    // Rest of your component code...
  
    onButtonClick(buttonId: string) {
      if (this.clickedButton === buttonId) {
        this.clickedButton = null; // Unselect the button if it's already clicked
      } else {
        this.clickedButton = buttonId; // Set the clickedButton to the current buttonId
      }
    }
  
    isButtonClicked(buttonId: string): boolean {
      return this.clickedButton === buttonId;
    }

  }
  