import type RGBModel from '../models/rgb'

export function get_color_from_percentage(percentage: number): RGBModel {
    // https://www.w3schools.com/colors/colors_picker.asp

    if (percentage < 10) {
        return <RGBModel>{
            // rgb(128, 0, 0)
            red: 128,
            green: 0,
            blue: 0
        };
    } else if (percentage < 20) {
        return <RGBModel>{
            // rgb(204, 0, 0)
            red: 204,
            green: 0,
            blue: 0
        };
    } else if (percentage < 30) {
        return <RGBModel>{
            // rgb(255, 80, 80)
            red: 255,
            green: 80,
            blue: 80
        };
    } else if (percentage < 40) {
        return <RGBModel>{
            // rgb(255, 204, 102)
            red: 255,
            green: 204,
            blue: 102
        };
    } else if (percentage < 50) {
        return <RGBModel>{
            // rgb(255, 255, 0)
            red: 255,
            green: 255,
            blue: 100
        };
    } else if (percentage < 60) {
        return <RGBModel>{
            // rgb(102, 255, 102)
            red: 102,
            green: 255,
            blue: 102
        };
    } else if (percentage < 70) {
        return <RGBModel>{
            // rgb(102, 255, 204)
            red: 102,
            green: 255,
            blue: 204
        };
    } else if (percentage < 80) {
        return <RGBModel>{
            // rgb(0, 204, 255)
            red: 0,
            green: 204,
            blue: 255
        };
    } else if (percentage < 90) {
        return <RGBModel>{
            // rgb(51, 102, 255)
            red: 51,
            green: 102,
            blue: 255
        };
    }
    return <RGBModel>{
        // rgb(51, 51, 153)
        red: 51,
        green: 51,
        blue: 153
    };
}