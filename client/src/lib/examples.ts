import type Arguments from "./models/arguments";

export type Example = {
	name: string;
	text: string
	arguments?: Partial<Arguments>
}
export const examples: Example[] = [
	{
		name: "Simple",
		text: "#\n# ultra simple demo - krpsim\n#\n# stock      name:quantity\neuro:10\n#\n# process   name:(need1:qty1;need2:qty2;[...]):(result1:qty1;result2:qty2;[...]):delay\n#\nachat_materiel:(euro:8):(materiel:1):10\nrealisation_produit:(materiel:1):(produit:1):30\nlivraison:(produit:1):(client_content:1):20\n#\n# optimize time for no process possible (eating stock, produce all possible),\n# or maximize some products over a long delay\n# optimize:(time|stock1;time|stock2;...)\n#\noptimize:(time;client_content)\n#\n",
		arguments: {
			max_generations: 5,
			max_cycle: 60,
			time_limit: 5000,
			population_size: 1,
		}
	},
	{
		name: "Steak",
		text: "#\n# steak demo - krpsim\n#\n# stock      name:quantity\nsteak_cru:3\npoele:1\n#\n# process   name:(need1:qty1;need2:qty2;[...]):(result1:qty1;result2:qty2;[...]):delay\n#\ncuisson_1:(steak_cru:2;poele:1):(steak_mi_cuit:2;poele:1):10\ncuisson_2:(steak_mi_cuit:2;poele:1):(steak_cuit:2;poele:1):10\ncuisson_3:(steak_cru:1;steak_mi_cuit:1;poele:1):(steak_mi_cuit:1;steak_cuit:1;poele:1):10\ncuisson_4:(steak_cru:1;poele:1):(steak_mi_cuit:1;poele:1):10\ncuisson_5:(steak_mi_cuit:1;poele:1):(steak_cuit:1;poele:1):10\n#\n# optimize time for 0 stock and no process possible,\n# or maximize some products over a long delay\n# optimize:(stock1;stock2;...)\n#\noptimize:(time;steak_cuit)\n#\n",
		arguments: {
			max_generations: 5,
			max_cycle: 50,
			time_limit: 5000,
			population_size: 1,
		}
	},
	{
		name: "Pomme",
		text: "#\n#  krpsim tarte aux pommes\n#\nfour:10\neuro:10000\n#\nbuy_pomme:(euro:100):(pomme:700):200\nbuy_citron:(euro:100):(citron:400):200\nbuy_oeuf:(euro:100):(oeuf:100):200\nbuy_farine:(euro:100):(farine:800):200\nbuy_beurre:(euro:100):(beurre:2000):200\nbuy_lait:(euro:100):(lait:2000):200\n#\nseparation_oeuf:(oeuf:1):(jaune_oeuf:1;blanc_oeuf:1):2\nreunion_oeuf:(jaune_oeuf:1;blanc_oeuf:1):(oeuf:1):1\ndo_pate_sablee:(oeuf:5;farine:100;beurre:4;lait:5):(pate_sablee:300;blanc_oeuf:3):300\ndo_pate_feuilletee:(oeuf:3;farine:200;beurre:10;lait:2):(pate_feuilletee:100):800\ndo_tarte_citron:(pate_feuilletee:100;citron:50;blanc_oeuf:5;four:1):(tarte_citron:5;four:1):60\ndo_tarte_pomme:(pate_sablee:100;pomme:30;four:1):(tarte_pomme:8;four:1):50\ndo_flan:(jaune_oeuf:10;lait:4;four:1):(flan:5;four:1):300\ndo_boite:(tarte_citron:3;tarte_pomme:7;flan:1;euro:30):(boite:1):1\nvente_boite:(boite:100):(euro:55000):30\nvente_tarte_pomme:(tarte_pomme:10):(euro:100):30\nvente_tarte_citron:(tarte_citron:10):(euro:200):30\nvente_flan:(flan:10):(euro:300):30\n#do_benef:(euro:1):(benefice:1):0\n#\n#\n#optimize:(benefice)\noptimize:(euro)\n#\n",
		arguments: {
			max_generations: 10,
			population_size: 4,
			max_cycle: 50000,
			time_limit: 600000,
			elitism_amount: 1,
		}
	},
	{
		name: "Inception",
		text: "#\n# inception demo - krpsim\n#\n# stock      name:quantity\nclock:1\n#\n# process   name:(need1:qty1;need2:qty2;[...]):(result1:qty1;result2:qty2;[...]):delay\n#\nmake_sec:(clock:1):(clock:1;second:1):1\nmake_minute:(second:60):(minute:1):6\nmake_hour:(minute:60):(hour:1):36\nmake_day:(hour:24):(day:1):86\nmake_year:(day:365):(year:1):365\nstart_dream:(minute:1;clock:1):(dream:1):60\nstart_dream_2:(minute:1;dream:1):(dream:2):60\ndream_minute:(second:1;dream:1):(minute:1;dream:1):1\ndream_hour:(second:1;dream:2):(hour:1;dream:2):1\ndream_day:(second:1;dream:3):(day:1;dream:3):1\nend_dream:(dream:3):(clock:1):60\n#\n# optimize time for no process possible (eating stock, produce all possible),\n# or maximize some products over a long delay\n# optimize:(time|stock1;time|stock2;...)\n#\noptimize:(year)\n#\n",
		arguments: {
			max_generations: 10,
			population_size: 4,
			max_cycle: 50000,
			time_limit: 600000,
			elitism_amount: 1,
		}
	},
	{
		name: "Ikea",
		text: "#\n# ikea demo - krpsim\n#\n# stock      name:quantity\nplanche:7\n#\n# process   name:(need1:qty1;need2:qty2;[...]):(result1:qty1;result2:qty2;[...]):delay\n#\ndo_montant:(planche:1):(montant:1):15\ndo_fond:(planche:2):(fond:1):20\ndo_etagere:(planche:1):(etagere:1):10\ndo_armoire_ikea:(montant:2;fond:1;etagere:3):(armoire:1):30\n#\n# optimize time for 0 stock and no process possible,\n# or maximize some products over a long delay\n# optimize:(stock1;stock2;...)\n#\noptimize:(time;armoire)\n#\n",
		arguments: {
			max_generations: 5,
			max_cycle: 50,
			time_limit: 5000,
			population_size: 1,
		}
	},
	{
		name: "Recre",
		text: "#\n# recre demo - krpsim\n#\n# stock      name:quantity\nbonbon:10\nmoi:1\n#\n# process   name:(need1:qty1;need2:qty2;[...]):(result1:qty1;result2:qty2;[...]):delay\n#\nmanger:(bonbon:1)::10\njouer_a_la_marelle:(bonbon:5;moi:1):(moi:1;marelle:1):20\nparier_avec_un_copain:(bonbon:2;moi:1):(moi:1;bonbon:3):10\nparier_avec_un_autre_copain:(moi:1;bonbon:2):(moi:1;bonbon:1):10\nse_battre_dans_la_cours:(moi:1):(moi:1;bonbon:1):50\n#\n# optimize time for no process possible (eating stock, produce all possible),\n# or maximize some products over a long delay\n# optimize:(time|stock1;time|stock2;...)\n#\noptimize:(marelle)\n#\n",
		arguments: {
			max_generations: 10,
			population_size: 4,
			max_cycle: 5000,
			time_limit: 600000,
			elitism_amount: 1,
		}
	},
	// {
	// 	name: "Toothpick",
	// 	text: "# ============================= RESOURCES\n# ---- primitive material\nwood:10\npackaging:10\n\n# ---- constante\nmachine:1\nemployee:2\n\n# ---- manufactured\ntoothpick:0\n\n# ---- salable\nfinal_product:0\n\n# ---- goal\ndollars:10\n\n# ============================= RENTABILITY\n# invest 2 dollars in primitive material product bring back 3 dollars:\n# - 2 dollars =\n# - 1 dollars + 1 dollars =\n# - 1 wood + 1 packaging =\n# - 100 toothpick + 1 packaging =\n# - 1 final product =\n# - 3 dollars\n\n# ============================= PROCESS\n# ---- buy primitive material\nbuy_wood:(dollars:20;employee:1):(wood:10;employee:1):5\nbuy_packaging:(dollars:10;employee:1):(packaging:10;employee:1):2\n\n# ---- manufacture\nmanufacture:(wood:10;employee:1;machine:1):(toothpick:1000;employee:1;machine:1):10\n# ---- sale\nfinal_product:(toothpick:100;packaging:1;employee:1):(final_product:1;employee:1):1\n# ---- goal\nsale:(final_product:10;employee:1):(dollars:20;employee:1):3\n\noptimize:(time;dollars)\n"
	// },
]
export default examples
