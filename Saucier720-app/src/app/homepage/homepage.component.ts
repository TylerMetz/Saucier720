import { Component, ViewChild, ElementRef, Renderer2 } from '@angular/core';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})

export class HomePageComponent{

  // used to determine which component to display
  noUserActive: boolean;
  userActive: boolean;

  constructor(private appComponent: AppComponent, private renderer: Renderer2) {
    this.noUserActive = !appComponent.getAuthService().isLoggedIn();
    this.userActive = appComponent.getAuthService().isLoggedIn();
  }

  buttonsVisible: boolean = false;
  isComponentReady = false;

  positionValue: string = 'absolute'; // Default position


  onGenerationComplete(generationComplete: boolean) {
    if (generationComplete) {
      // The logo generation is complete, you can now trigger the buttons generation or their appearance.
      // For example, you can set a flag to control their visibility:
      this.buttonsVisible = true;
      this.isComponentReady = true;
    }
  }

  buttonMovement(cardExpanded: boolean) {
    if (cardExpanded) {
      // change button to be relative
      this.positionValue = 'relative';
    }
    else {
      // change button to be absolute
      setTimeout(() => {
        this.positionValue = 'absolute';
      }, 650); // to wait for recipe card to collapse past the point of the buttons' location
    }
  }

activeButton: string = '';
hoveredButton: string = '';

isButtonClicked(buttonName: string): boolean {
  return this.activeButton === buttonName;
}

onButtonClick(buttonName: string): void {
  this.activeButton = this.activeButton === buttonName ? '' : buttonName;
}

onButtonHover(buttonName: string): void {
  this.hoveredButton = buttonName;
}

onButtonLeave(buttonName: string): void {
  this.hoveredButton = '';
}

isButtonHovered(buttonName: string): boolean {
  return this.hoveredButton === buttonName;
}

}
