import { Component } from '@angular/core';


@Component({
    selector: 'app-landing-page',
    templateUrl: './landing-page.component.html',
    styleUrls: ['./landing-page.component.scss']
  })

  export class LandingPageComponent {
    buttonsVisible: boolean = false;

    onLogoGenerationComplete(generationComplete: boolean) {
      if (generationComplete) {
        // The logo generation is complete, you can now trigger the buttons generation or their appearance.
        // For example, you can set a flag to control their visibility:
        this.buttonsVisible = true;
      }
    }

    // Add a property to keep track of clicked buttons
  clickedButtons: Set<string> = new Set<string>();
  
  // Method to handle button clicks
  onButtonClick(buttonId: string) {
    if (this.clickedButtons.has(buttonId)) {
      this.clickedButtons.delete(buttonId);
    } else {
      this.clickedButtons.add(buttonId);
    }
  }

  // Method to check if a button is clicked
  isButtonClicked(buttonId: string): boolean {
    return this.clickedButtons.has(buttonId);
  }
  }
  