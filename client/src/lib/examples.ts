import type Arguments from "./models/arguments";

export type Example = {
	name: string;
	text: string
	arguments?: Partial<Arguments>
}
export const examples: Example[] = [
	{
		name: "Furniture",
		text: "master:1\n\nbuy_slave:(master:1;gold:5):(master:1;slave:1):20\nbuy_wood:(master:1;gold:10):(master:1;wood:50):10\nbuy_paint:(master:1;gold:1):(master:1;paint:3):5\nbuy_planck:(master:1;gold:10):(master:1;planck:1200):20\n\nstill_wood_with_master:(master:1):(master:1;wood:3):30\nstill_wood_with_slave:(slave:1):(slave:1;wood:3):20\n\nmake_planck:(wood:1):(planck:10):20\n\ncraft_furniture:(slave:1;planck:50):(slave:1;furniture:1):50\ncraft_colored_furniture:(slave:1;planck:50;paint:2):(slave:1;colored_furniture:1):80\n\nsell_wood_with_master:(master:1;wood:7):(master:1;gold:1):15\nsell_wood_with_slave:(slave:1;wood:8):(slave:1;gold:1):5\nsell_planck_with_master:(master:1;planck:95):(master:1;gold:1):10\nsell_planck_with_slave:(slave:1;planck:95):(slave:1;gold:1):10\nsell_furniture_1_with_master:(master:1;furniture:1):(master:1;gold:7):5\nsell_furniture_2_with_master:(master:1;furniture:2):(master:1;gold:15):15\nsell_colored_furniture_with_master:(master:1;colored_furniture:1):(master:1;gold:20):5\n\noptimize:(gold)",
		arguments: {
			max_generations: 8,
			max_cycle: 3000,
			time_limit: 30000,
			population_size: 8,
			max_depth: 7,
		}
	},
	{
		name: "Bakery",
		text: "oven:10\ndollar:10000\n\nbuy_apple:(dollar:100):(apple:700):200\nbuy_lemon:(dollar:100):(lemon:400):200\nbuy_egg:(dollar:100):(egg:100):200\nbuy_flour:(dollar:100):(flour:800):200\nbuy_butter:(dollar:100):(butter:2000):200\nbuy_milk:(dollar:100):(milk:2000):200\n\nseparation_egg:(egg:1):(egg_fauna:1;egg_white:1):2\nmerge_egg_parts:(egg_fauna:1;egg_white:1):(egg:1):1\ndo_shortbread:(egg:5;flour:100;butter:4;milk:5):(shortbread:300;egg_white:3):300\ndo_puff_pastry:(egg:3;flour:200;butter:10;milk:2):(puff_pastry:100):800\n\ndo_lemon_pie:(puff_pastry:100;lemon:50;egg_white:5;oven:1):(lemon_pie:5;oven:1):60\ndo_apple_pie:(shortbread:100;apple:30;oven:1):(apple_pie:8;oven:1):50\ndo_flan:(egg_fauna:10;milk:4;oven:1):(flan:5;oven:1):300\ndo_box:(lemon_pie:3;apple_pie:7;flan:1;dollar:30):(box:1):1\n\nsell_box:(box:100):(dollar:55000):30\nsell_apple_pie:(apple_pie:10):(dollar:100):30\nsell_lemon_pie:(lemon_pie:10):(dollar:200):30\nsell_flan:(flan:10):(dollar:300):30\n\noptimize:(dollar)",
		arguments: {
			max_generations: 5,
			population_size: 6,
			max_cycle: 50000,
			time_limit: 600000,
			elitism_amount: 1,
			max_depth: 6,
		}
	},
	{
		name: "Inception",
		text: "clock:1\n\nmake_sec:(clock:1):(clock:1;second:1):1\nmake_minute:(second:60):(minute:1):6\nmake_hour:(minute:60):(hour:1):36\nmake_day:(hour:24):(day:1):86\nmake_year:(day:365):(year:1):365\nstart_dream:(minute:1;clock:1):(dream:1):60\nstart_dream_2:(minute:1;dream:1):(dream:2):60\ndream_minute:(second:1;dream:1):(minute:1;dream:1):1\ndream_hour:(second:1;dream:2):(hour:1;dream:2):1\ndream_day:(second:1;dream:3):(day:1;dream:3):1\nend_dream:(dream:3):(clock:1):60\noptimize:(year)",
		arguments: {
			max_generations: 3,
			population_size: 5,
			max_cycle: 50000,
			time_limit: 600000,
			elitism_amount: 1,
			max_depth: 6,
		}
	},
	{
		name: "Farm",
		text: "farmer:2\ntired_farmer:2\n\ntractor:1\nvan:1\nessence:100\n\ncow:10\nsheep:20\nchicken:300\n\nready_cow:10\nready_sheep:20\nready_chicken:30\n\nwheat_field:2\nsown_wheat_field:1\nready_wheat_field:0\ncorn_field:0\nsown_corn_field:3\nready_corn_field:0\n\nwheat:1000\ncorn:1400\nmilk:1000\nwool:200\negg:300\ndollar:500\n\nbuy_van:(farmer:1;dollar:5000):(tired_farmer:1;van:1):8\nbuy_tractor:(farmer:1;dollar:2000):(tired_farmer:1;tractor:1):6\nbuy_essence:(farmer:1;dollar:100):(farmer:1;essence:100):1\n\nfarmer_rest:(tired_farmer:1):(farmer:1):12\ncow_rest:(cow:1):(ready_cow:1):24\nsheep_rest:(sheep:1):(ready_sheep:1):700\nchicken_rest:(chicken:1):(ready_chicken:1):24\n\nsow_wheat_field:(farmer:1;tractor:1;wheat_field:1;essence:60):(tired_farmer:1;tractor:1;sown_wheat_field:1):8\nsow_corn_field:(farmer:1;tractor:1;corn_field:1;essence:45):(tired_farmer:1;tractor:1;sown_corn_field:1):7\n\ngrow_wheat_field:(sown_wheat_field:1):(ready_wheat_field:1):900\ngrow_corn_field:(sown_corn_field:1):(ready_corn_field:1):800\n\nharvest_wheat_field:(farmer:2;tractor:1;essence:120;ready_wheat_field:1):(tired_farmer:2;tractor:1;wheat_field:1;weath:1000):10\nharvest_corn_field:(farmer:2;tractor:1;essence:90;ready_corn_field:1):(tired_farmer:2;tractor:1;corn_field:1;corn:1400):9\n\nmilk_cow:(farmer:1;ready_cow:10):(tired_farmer:1;cow:10;milk:200):6\nshear_sheep:(farmer:1;ready_sheep:20):(tired_farmer:1;sheep:20;wool:100):6\npick_up_egg:(farmer:1;ready_chicken:150):(tired_farmer:1;chicken:150;egg:150):6\n\nsell_wheat:(farmer:1;van:1;essence:15;wheat:1000):(tired_farmer:1;van:1;dollar:1000):8\nsell_corn:(farmer:1;van:1;essence:15;corn:1400):(tired_farmer:1;van:1;dollar:800):8\nsell_milk:(farmer:1;van:1;essence:10;milk:400):(tired_farmer:1;van:1;dollar:400):6\nsell_wool:(farmer:1;van:1;essence:10;wool:200):(tired_farmer:1;van:1;dollar:1200):6\nsell_egg:(farmer:1;van:1;essence:10;egg:600):(tired_farmer:1;van:1;dollar:300):6\n\noptimize:(dollar)",
		arguments: {
			max_generations: 6,
			population_size: 7,
			max_cycle: 5000,
			time_limit: 600000,
			elitism_amount: 1,
			max_depth: 7,
		}
	},
	{
		name: "Furniture",
		text: "master:1\n\nbuy_slave:(master:1;gold:5):(master:1;slave:1):20\nbuy_wood:(master:1;gold:10):(master:1;wood:50):10\nbuy_paint:(master:1;gold:1):(master:1;paint:3):5\nbuy_planck:(master:1;gold:10):(master:1;planck:1200):20\n\nstill_wood_with_master:(master:1):(master:1;wood:3):30\nstill_wood_with_slave:(slave:1):(slave:1;wood:3):20\n\nmake_planck:(wood:1):(planck:10):20\n\ncraft_furniture:(slave:1;planck:50):(slave:1;furniture:1):50\ncraft_colored_furniture:(slave:1;planck:50;paint:2):(slave:1;colored_furniture:1):80\n\nsell_wood_with_master:(master:1;wood:7):(master:1;gold:1):15\nsell_wood_with_slave:(slave:1;wood:8):(slave:1;gold:1):5\nsell_planck_with_master:(master:1;planck:95):(master:1;gold:1):10\nsell_planck_with_slave:(slave:1;planck:95):(slave:1;gold:1):10\nsell_furniture_1_with_master:(master:1;furniture:1):(master:1;gold:7):5\nsell_furniture_2_with_master:(master:1;furniture:2):(master:1;gold:15):15\nsell_colored_furniture_with_master:(master:1;colored_furniture:1):(master:1;gold:20):5\n\noptimize:(gold)",
		arguments: {
			max_generations: 6,
			population_size: 8,
			max_cycle: 2000,
			time_limit: 600000,
			elitism_amount: 1,
			max_depth: 7,
		}
	},
	{
		name: "House",
		text: "worker:3\ngold:50\n\nbuy_wood:(worker:1;gold:10):(worker:1;wood:10):1\nbuy_stone:(worker:1;gold:10):(worker:1;stone:10):1\n\nbuy_wood_planck:(worker:1;gold:10):(worker:1;wood_planck:100):1\nbuy_stone_brick:(worker:1;gold:10):(worker:1;stone_brick:100):1\n\nmake_wood_planck:(worker:1;wood:5):(worker:1;wood_planck:50):12\nmake_stone_brick:(worker:2;stone:5):(worker:2;stone_brick:50):12\n\nbuild_house:(worker:2;wood_planck:100;stone_brick:100):(worker:2;house:1):48\nsell_house:(worker:1;house:1):(worker:1;gold:150):24\n\noptimize:(gold)",
		arguments: {
			max_generations: 10,
			population_size: 50,
			max_cycle: 300,
			time_limit: 600000,
			elitism_amount: 1,
			max_depth: 4,
		}
	},
]
export default examples
