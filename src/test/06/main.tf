provider "azurerm" {
  subscription_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}

variable "location" {
  type    = string
  default = "japaneast"
}

variable "vnet_cidr" {
  type    = string
  default = "10.0.0.0/16"
}

resource "azurerm_resource_group" "rg" {
  name     = "my-resource-group"
  location = "japaneast"
}

resource "azurerm_virtual_network" "vnet" {
  name                = "my-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = "japaneast"
  resource_group_name = "my-resource-group"
}
