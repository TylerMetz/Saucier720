import { Component } from '@angular/core';
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

  constructor(private appComponent: AppComponent) {
    this.noUserActive = !appComponent.getAuthService().isLoggedIn();
    this.userActive = appComponent.getAuthService().isLoggedIn();
  }

  buttonsVisible: boolean = false;
  isComponentReady = false;


  onGenerationComplete(generationComplete: boolean) {
    if (generationComplete) {
      // The logo generation is complete, you can now trigger the buttons generation or their appearance.
      // For example, you can set a flag to control their visibility:
      this.buttonsVisible = true;
      this.isComponentReady = true;
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

getTopMargin(): number {
  // Define your margin values based on conditions
  if (this.noUserActive) {
    return 15; // Adjust the value as needed
  } else if (this.userActive) {
    return 10.25; // Adjust the value as needed
  } else {
    return 0; // Default margin value
  }
}

}
