import { Component } from '@angular/core';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})

export class HomePageComponent{
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
