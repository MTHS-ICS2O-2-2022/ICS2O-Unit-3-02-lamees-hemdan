// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Lamees Hemdan
// Created on: March 2023
// This file contains the JS functions for index.html

// This function calculates the area of a trapezoid

function calculateVolume () {
 // input
  const length = parseFloat(document.getElementById('length').value)
  const width = parseFloat(document.getElementById('width').value)
  const height = parseFloat(document.getElementById('height').value)


 // process
  const volume = (length * height * width) / 3


 // output
  document.getElementById('volume').innerHTML = 'The Volume is ' + volume + ' mm³'
}
