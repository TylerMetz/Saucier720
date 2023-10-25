import { Component, OnInit } from '@angular/core';
import { Ingredient, List } from '../core/interfaces/ingredient';
import { ListService } from '../core/services/list/list.service';
import { HttpClient } from '@angular/common/http';
import { HttpEvent, HttpEventType } from "@angular/common/http"
import { lastValueFrom } from 'rxjs';
import { AuthService } from '../core/services/Auth/auth.service';
import { GetListRequest, PostListRequest, UpdateListRequest } from '../core/interfaces/types';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {
  name: string = '';
  quantity: number = 0;

  list: List = {
    Ingredients: []
  }

  constructor(private listService: ListService, private http: HttpClient, private authService: AuthService) { }

  async ngOnInit() {
    await this.populateList();
  }

  public async populateList(): Promise<void> {
    const request: GetListRequest = {
      UserName: this.authService.getUsername(),
    };
    console.log('request: ', request);
    this.listService.getList(request.UserName).subscribe({
      next: (response: any) => {
        console.log('GetListResponse ', response);
        this.list = response.List;
        console.log('List Updated: ', this.list);
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  async postListItem() {
    if(!this.name){
      return; 
    }
    const request: PostListRequest = {
      UserName: this.authService.getUsername(),
      Ingredient: {
        Name: this.name,
        FoodType: '',
        SaleDetails: '',
        Quantity: this.quantity,
      }
    };
    this.listService.postListItem(request).subscribe({
      next: (response: any) => {
        console.log('New Shopping List Item Posted: ', response)
        this.name = '';
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  async updateList(){
    const zeroQuantityItems = this.list.Ingredients.filter((item: Ingredient) => item.Quantity === 0); 
    this.list.Ingredients = this.list.Ingredients.filter((item: Ingredient) => item.Quantity !== 0);
    const request: UpdateListRequest = { 
      UserName: this.authService.getUsername(),
      List: this.list,
      ItemsToDelete: zeroQuantityItems,
    };
    const response = await lastValueFrom(this.listService.updateList(request));
    console.log('UpdateListResponse: ', response);
    this.populateList();
  }
}
