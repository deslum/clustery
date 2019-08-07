package dbs

import "encoding/json"

const jsonPointsData = "[[30.258387, 59.951557], [30.434124, 60.029499], [37.642517, 55.828777], [55.956112, 54.702778], [30.163561, 59.809929], [39.822922, 57.672379], [30.222057, 59.987171], [39.838715, 57.673241], [39.816383, 57.683563], [30.304995, 59.906403], [30.240265, 59.951828], [55.956112, 54.702778], [30.321133, 59.808334], [30.398251, 59.798973], [30.398251, 59.798973], [37.928783, 59.122658], [30.215363, 59.985455], [30.269926, 59.95079], [30.294989, 59.808968], [30.269926, 59.95079], [34.348171, 61.781609], [37.722153, 44.754757], [85.173615, 52.511944], [30.305201, 59.906319], [30.274799, 59.905045], [30.10704, 59.848518], [30.10704, 59.848518], [30.093334, 59.854534], [49.472778, 53.486389], [37.917946, 59.089333], [30.259613, 59.945427], [30.301603, 59.911541], [30.10704, 59.848518], [37.917946, 59.089333], [44.032108, 56.300205], [34.394905, 61.784111], [30.446316, 59.908722], [30.48535, 59.93512], [30.143106, 59.823051], [82.921669, 55.118889], [30.223917, 59.984699], [30.240755, 59.952072], [30.24472, 59.955975], [30.366283, 59.81361], [30.24472, 59.955975], [30.626213, 59.726067], [37.628971, 55.824425], [30.241648, 59.956799], [30.418478, 60.027519], [39.846497, 59.186054], [30.111698, 59.847164], [55.956112, 54.702778], [37.444386, 55.82655], [30.434814, 60.027538], [29.892851, 59.795109], [44.961666, 53.224445], [49.356831, 53.551579], [30.233097, 59.87418], [30.311964, 59.812389], [38.989265, 55.810677], [86.618256, 53.869011], [44.082855, 56.300865], [30.240265, 59.951828], [82.942223, 55.115253], [30.442141, 59.910454], [30.319765, 60.044281], [82.942223, 55.115253], [49.361942, 53.544167], [30.10704, 59.848518], [30.24472, 59.955975], [34.365894, 61.773148], [30.366207, 59.734398], [30.433374, 60.032806], [37.962112, 59.125305], [30.222057, 59.987171], [30.442352, 59.906124], [30.228682, 59.961269], [30.294121, 59.913307], [30.22628, 59.984489], [37.917946, 59.089333], [30.417412, 60.036903], [30.475889, 59.946159], [30.316261, 60.040699], [29.880516, 59.670269], [30.22628, 59.984489], [37.900555, 59.120556], [37.770718, 55.69952], [37.396965, 55.798645], [37.887531, 55.75211], [30.304995, 59.906403], [30.30105, 59.906067], [30.163561, 59.809929], [39.786743, 54.623516], [82.947525, 55.108124], [30.301603, 59.909935], [30.228937, 59.955486], [30.315216, 59.810055], [30.494091, 59.933277], [30.412436, 60.019444], [30.43158, 59.946812], [30.304995, 59.906403], [30.383959, 59.814262], [49.354168, 53.547222], [34.354527, 61.795582], [30.256298, 59.958683], [30.316261, 60.040699], [30.491566, 59.945103], [30.464399, 60.047668], [30.311195, 59.961575], [30.24472, 59.955975], [30.24472, 59.955975], [34.357529, 61.795834], [44.0658, 56.300865], [34.357529, 61.795834], [30.154856, 59.809929], [37.905693, 59.090752], [44.0658, 56.300865], [39.898235, 59.219349], [37.923332, 59.085556], [30.155834, 59.809006], [30.492496, 59.937935], [29.990747, 59.820244], [85.175552, 52.529167], [30.323803, 59.808727], [49.356949, 53.546619], [30.32514, 59.809891], [30.317495, 60.053444], [30.465143, 59.960094], [37.9851, 55.798199], [82.942223, 55.111389], [37.917946, 59.089333], [30.319202, 60.045364], [30.433372, 59.974979], [49.319721, 53.548889], [30.311195, 59.961575], [30.678865, 59.726051], [37.883728, 55.787594], [49.458611, 53.484165], [44.0658, 56.300865], [39.943333, 57.656326], [37.977474, 55.80064], [39.943333, 57.656326], [39.788334, 54.620098], [30.301603, 59.911541], [30.347841, 59.789471], [30.495152, 59.936497], [37.917946, 59.089333], [30.319202, 60.045364], [30.378305, 59.809898], [30.301603, 59.911541], [85.178482, 52.517792], [85.178482, 52.517792], [39.943333, 57.656326], [37.923306, 59.09375], [73.459656, 61.261547], [30.398174, 59.801132], [37.917946, 59.089333], [37.917946, 59.089333], [30.32514, 59.809891], [30.102442, 59.755318], [73.459656, 61.261547], [30.315216, 59.810055], [30.305201, 59.906319], [30.434814, 60.027538], [30.306477, 59.904129], [44.048553, 56.288578], [30.226967, 59.838863], [73.449112, 61.25729], [39.967579, 57.615368], [46.001667, 51.598331], [30.316261, 60.040699], [30.494766, 59.947014], [30.427782, 60.030994], [30.313948, 60.049637], [39.838715, 57.673241], [30.433374, 60.032806], [37.933613, 59.160557], [37.906082, 59.127373], [30.323803, 59.808727], [30.281393, 59.916348], [30.665346, 59.952682], [37.37714, 59.187527], [37.907055, 59.123611], [86.635559, 53.861958], [55.956112, 54.702778], [39.838715, 57.673241], [30.245201, 59.959896], [30.434124, 60.029499], [30.32514, 59.809891], [30.427782, 60.030994], [30.315357, 59.807198], [39.842876, 59.19659], [30.398251, 59.798973], [30.451073, 59.919037], [55.119999, 51.842499], [29.863771, 59.867134], [30.244602, 59.955917], [30.107775, 59.818859], [30.434814, 60.027538], [37.9851, 55.798199], [30.46216, 59.899651], [39.843777, 59.193974], [39.846497, 59.186054], [34.357529, 61.795834], [46.001667, 51.598331], [30.434814, 60.027538], [30.304995, 59.906403], [37.469891, 55.726955], [30.326496, 59.811073], [30.28503, 59.912971], [30.304995, 59.906403], [37.334568, 55.699932], [34.34557, 61.797569], [39.91386, 57.661652], [30.432152, 59.946026], [30.448624, 59.910351], [30.24472, 59.955975], [36.57172, 50.565109], [39.837872, 57.668083], [30.435919, 59.945473], [30.305201, 59.906319], [34.354527, 61.795582], [37.9589, 55.793499], [30.319431, 60.044483], [30.10704, 59.848518], [30.37072, 60.009029], [34.357529, 61.795834], [30.093334, 59.85556], [37.917946, 59.089333], [34.354527, 61.795582], [30.311073, 60.047928], [30.4863, 59.929993], [34.354527, 61.795582], [30.17742, 59.807434], [30.412264, 59.795349], [49.321388, 53.533611], [30.304995, 59.906403], [37.911888, 59.09864], [30.433374, 60.032806], [30.458305, 59.914459], [37.93261, 59.131889], [34.357529, 61.795834], [37.859818, 59.129833], [30.24472, 59.955975], [34.354527, 61.795582], [34.350929, 61.786236], [30.301603, 59.911541], [30.530924, 59.771492], [30.344059, 60.056522], [30.17696, 59.825451], [30.093334, 59.854534], [37.424526, 55.726292], [30.215363, 59.985455], [47.258827, 47.393539], [30.26918, 59.931633], [44.961666, 53.224445], [37.884361, 59.133583], [30.181974, 59.82613], [37.444386, 55.82655], [30.753729, 59.635494], [30.323803, 59.808727], [30.326645, 59.917656], [39.179169, 51.717499], [30.417341, 60.036743], [43.103508, 44.191944], [30.398174, 59.801132], [44.961666, 53.224445], [30.155834, 59.809006], [30.434814, 60.027538], [30.093334, 59.854534], [30.222057, 59.987171], [85.183609, 52.52639], [46.003017, 51.599789], [30.24472, 59.955975], [30.28533, 59.926342], [39.232166, 51.69875], [30.222057, 59.987171], [39.837654, 59.211777], [37.905693, 59.090752], [37.982121, 55.798565], [73.439758, 61.250549], [39.79274, 54.613075], [30.259916, 59.946293], [49.354168, 53.547222], [82.942223, 55.111389], [30.111698, 59.847164], [30.323803, 59.808727], [30.301088, 59.907742], [30.111698, 59.847164], [30.24472, 59.955975], [30.222057, 59.987171], [39.179169, 51.717499], [30.418261, 60.008503], [49.354168, 53.547222], [30.093334, 59.853798], [30.222057, 59.987171], [34.357529, 61.795834], [39.837872, 57.668083], [34.354527, 61.795582], [30.304995, 59.906403], [30.284208, 59.732933], [30.03228, 59.855011], [37.966045, 55.758251], [30.439301, 59.908752], [30.43483, 59.94669], [37.485554, 55.731613], [30.444202, 59.912117], [30.479795, 59.957111], [45.843334, 53.174168], [30.24472, 59.955975], [37.900555, 59.120556], [37.918869, 59.080795], [30.427782, 60.030994], [37.863804, 55.767727], [30.244759, 59.955982], [39.935604, 57.66478], [30.313948, 60.049637], [30.614849, 59.731327], [30.222057, 59.987171], [30.324594, 59.886959], [30.321112, 59.808956], [30.344084, 59.985577], [30.319431, 60.044483], [30.110039, 59.847805], [30.614849, 59.731327], [38.142944, 59.042419], [30.259567, 59.878056], [30.243324, 59.957798], [29.15855, 60.375378], [34.354527, 61.795582], [37.973999, 55.793999], [30.678865, 59.726051], [30.222057, 59.987171], [30.222057, 59.987171], [30.24472, 59.955975], [30.495152, 59.936497], [30.093334, 59.853798], [30.111698, 59.847164], [30.162861, 59.807167], [34.337158, 61.794823], [30.21718, 59.98576], [49.473701, 53.498291], [30.315216, 59.810055], [30.444202, 59.912117], [30.304995, 59.906403], [30.155834, 59.809006], [37.440723, 55.731659], [30.509666, 59.933533], [30.444202, 59.912117], [30.434814, 60.027538], [30.24472, 59.955975], [46.001667, 51.598331], [30.28503, 59.912971], [30.26158, 59.954937], [30.305841, 59.922596], [50.806854, 61.698402], [30.24472, 59.955975], [34.350983, 61.795288], [36.568531, 50.565289], [44.943333, 53.223331], [44.961666, 53.224445], [30.304998, 59.941277], [37.883728, 55.787594], [30.250736, 59.952724], [36.57172, 50.565109], [30.319431, 60.044483], [37.917946, 59.089333], [44.08337, 56.300865], [30.10704, 59.848518], [39.79274, 54.613075], [30.323803, 59.808727], [30.287428, 59.912277], [50.806854, 61.698402], [37.758884, 55.70232], [30.299744, 59.908398], [29.939953, 59.870312], [30.383055, 59.806625], [30.155834, 59.809006], [30.222057, 59.987171], [30.322643, 59.807716], [37.923306, 59.09375], [30.678865, 59.726051], [44.889999, 53.222221], [30.320679, 60.043739], [30.292629, 59.912815], [30.446951, 59.765621], [30.315216, 59.810055], [37.917946, 59.089333], [30.444202, 59.912117], [37.334568, 55.699932], [30.304995, 59.906403], [30.435919, 59.945473], [30.434814, 60.027538], [30.093334, 59.854534], [30.323803, 59.808727], [37.856239, 55.746342], [37.923306, 59.09375], [38.085228, 55.751438], [30.350834, 59.830704], [37.396965, 55.798645], [30.678865, 59.726051], [37.977474, 55.80064], [73.459656, 61.261547], [34.344975, 61.794956], [73.453178, 61.255978], [30.448603, 59.911942], [30.389116, 59.803574], [30.678865, 59.726051], [30.678865, 59.726051], [30.244602, 59.955917], [30.319431, 60.044483], [30.323215, 59.908554], [37.917946, 59.089333], [30.504477, 59.946659], [30.297989, 59.907173], [37.869804, 59.125111], [30.10704, 59.848518], [30.223917, 59.984699], [37.333935, 55.69841], [30.304995, 59.906403], [30.304995, 59.906403], [34.357529, 61.795834], [46.001667, 51.598331], [44.082855, 56.300865], [39.846497, 59.186054], [44.945, 53.233334], [30.398251, 59.798973], [30.316261, 60.040699], [30.433374, 60.032806], [85.173615, 52.511944], [30.222057, 59.987171], [44.082855, 56.300865], [37.846992, 55.745235], [30.45311, 59.903046], [37.628971, 55.824425], [30.442287, 60.040833], [30.106731, 59.819611], [37.785553, 55.750759], [30.215363, 59.985455], [30.155834, 59.809006], [39.83802, 57.668983], [30.480774, 59.904316], [37.910114, 59.094299], [30.222057, 59.987171], [30.315216, 59.810055], [30.216248, 59.985374], [30.262533, 59.953304], [30.215363, 59.985455], [34.361816, 61.773479], [37.917946, 59.089333], [44.943333, 53.223331], [30.217726, 59.953007], [30.433374, 60.032806], [30.412264, 59.795349], [30.456095, 59.906143], [37.758137, 55.700821], [30.223497, 59.987587], [30.280392, 59.901085], [30.244602, 59.955917], [30.323803, 59.808727], [30.266119, 59.952393], [30.327028, 59.781349], [30.398174, 59.801132], [37.759468, 55.702091], [30.319431, 60.044483], [49.472778, 53.486389], [37.923332, 59.085556], [30.491627, 59.939308], [30.435141, 59.96273], [30.222057, 59.987171], [37.917946, 59.089333], [37.9175, 55.7995], [38.036991, 55.808056], [37.751373, 55.706112], [85.17691, 52.515163], [85.183609, 52.52639], [30.323803, 59.808727], [30.442141, 59.910454], [30.399895, 59.799801], [30.222057, 59.987171], [30.200258, 59.849888], [30.222057, 59.987171], [30.315216, 59.810055], [49.473812, 53.48587], [30.268238, 59.800194], [30.433374, 60.032806], [30.405226, 59.966652], [30.24472, 59.955975], [30.495838, 59.934013], [39.837872, 57.668083], [30.494349, 59.937195], [30.434814, 60.027538], [30.309454, 59.904526], [30.093334, 59.854534], [30.162861, 59.807167], [37.9851, 55.798199], [30.433374, 60.032806], [37.985085, 55.798168], [30.3141, 60.047661], [30.434814, 60.027538], [30.184679, 59.862797], [37.730026, 44.755306], [44.952625, 53.220646], [30.163561, 59.809929], [30.155834, 59.809006], [39.792439, 54.614655], [30.432152, 59.946026], [34.33717, 61.797619], [30.316261, 60.040699], [30.306528, 59.905758], [30.269926, 59.95079], [30.383959, 59.814262], [39.788334, 54.620098], [37.92775, 59.090694], [30.244759, 59.955982], [37.444386, 55.82655], [30.446316, 59.908722], [39.788334, 54.620098], [30.216248, 59.985374], [30.448162, 59.916096], [30.155834, 59.809006], [30.10704, 59.848518], [30.456095, 59.906143], [30.323215, 59.908554], [30.111698, 59.847164], [44.082855, 56.300865], [30.093334, 59.853798], [30.317408, 60.045689], [39.935604, 57.66478], [37.396667, 55.798706], [30.295248, 59.808228], [30.162861, 59.807167], [30.443539, 59.912037], [39.943333, 57.656326], [44.019794, 56.316944], [30.572836, 59.693909], [30.222057, 59.987171], [37.628971, 55.824425], [30.111698, 59.847164], [39.172028, 51.715332], [37.765495, 55.703732], [37.759468, 55.702091], [30.428268, 59.975586], [37.758884, 55.70232], [30.515127, 59.851894], [34.357529, 61.795834], [30.096621, 59.797756], [30.614849, 59.731327], [30.315216, 59.810055], [30.222057, 59.987171], [37.959911, 59.104618], [30.216248, 59.985374], [30.265074, 59.883602], [49.354168, 53.547222], [30.222057, 59.987171], [30.323803, 59.808727], [30.397121, 60.021099], [37.92775, 59.090694], [30.309454, 59.904526], [37.628971, 55.824425], [30.398174, 59.801132], [30.444202, 59.912117], [49.472778, 53.486389], [30.444202, 59.912117], [37.628971, 55.824425], [37.628971, 55.824425], [30.10704, 59.848518], [30.163561, 59.809929], [49.354168, 53.547222], [30.256298, 59.958683], [30.497301, 59.934628], [34.354527, 61.795582], [34.354527, 61.795582], [30.457832, 59.975716], [37.884361, 59.133583], [49.360985, 53.544167], [37.930752, 59.089668], [30.448624, 59.910351], [49.361942, 53.544167], [30.446316, 59.908722], [30.446316, 59.908722], [30.430847, 59.953148], [30.495838, 59.934013], [73.454353, 61.259205], [30.434814, 60.027538], [73.454353, 61.259205], [30.321089, 59.833145], [30.324938, 59.809715], [30.572617, 59.956299], [30.319431, 60.044483], [30.270361, 59.933926], [30.372482, 59.933395], [37.923306, 59.09375], [30.243324, 59.957798], [55.956112, 54.702778], [27.919815, 60.594994], [39.833157, 54.613838], [30.111698, 59.847164], [73.449112, 61.25729], [30.111698, 59.847164], [73.460754, 61.261993], [37.730026, 44.755306], [30.446316, 59.908722], [30.262962, 59.834091], [85.177414, 52.515366], [30.176668, 59.82077], [30.428268, 59.975586], [30.20274, 59.839104], [49.478516, 53.490559], [30.40671, 59.803234], [30.022676, 59.705574], [30.48049, 59.940842], [30.297989, 59.907173], [29.892851, 59.795109], [30.43483, 59.94669], [30.308332, 60.065395], [30.304995, 59.906403], [37.9604, 55.787998], [30.515127, 59.851894], [85.17691, 52.515163], [30.266119, 59.952393], [56.064484, 54.828438], [37.396667, 55.798706], [37.883728, 55.787594], [44.952625, 53.220646], [37.396667, 55.798706], [30.215363, 59.985455], [30.389397, 59.804993], [30.227449, 59.956375], [37.765919, 55.679688], [30.37154, 60.081711], [44.943333, 53.223331], [30.319202, 60.045364], [37.986801, 55.793999], [29.943457, 59.875278], [30.234074, 59.94886], [37.856239, 55.746342], [55.956112, 54.702778], [37.923306, 59.09375], [37.910114, 59.094299], [30.155834, 59.809006], [34.357529, 61.795834], [50.739742, 61.819416], [30.396196, 59.941505], [30.315216, 59.810055], [30.2941, 59.925201], [38.081882, 55.776302], [28.554581, 60.936878], [30.252947, 59.832256], [43.927814, 56.275867], [30.40671, 59.803234], [30.106731, 59.819611], [30.296925, 59.866241], [30.181974, 59.82613], [30.163561, 59.809929], [49.323334, 53.520279], [30.223497, 59.987587], [30.308735, 59.824333], [30.216248, 59.985374], [30.497301, 59.934628], [30.398174, 59.801132], [30.162861, 59.807167], [30.222057, 59.987171], [30.25626, 59.877495], [30.31587, 60.046764], [49.354168, 53.547222], [30.3141, 60.047661], [30.481022, 59.915607], [34.357578, 61.785637], [34.357529, 61.795834], [30.497301, 59.934628], [37.917946, 59.089333], [55.996666, 54.696945], [86.618256, 53.869011], [30.222057, 59.987171], [37.918869, 59.080795], [82.921669, 55.118889], [30.231226, 59.946014], [44.08337, 56.300865], [34.433083, 61.75922], [34.357529, 61.795834], [45.968506, 51.611668], [45.019169, 53.192123], [30.301603, 59.911541], [37.911888, 59.09864], [73.457802, 61.260803], [49.354168, 53.547222], [39.838715, 57.673241], [37.92775, 59.090694], [30.402056, 59.800888], [73.452347, 61.24678], [30.586538, 59.739468], [30.245024, 59.956047], [37.968639, 59.12875], [30.266119, 59.952393], [37.92775, 59.090694], [30.427586, 60.014961], [30.256298, 59.958683], [30.24472, 59.955975], [30.678865, 59.726051], [30.495152, 59.936497], [30.678865, 59.726051], [73.459656, 61.261547], [30.266119, 59.952393], [44.082855, 56.300865], [86.618614, 53.867802], [30.156893, 59.808079], [30.320736, 60.040833], [30.243324, 59.957798], [34.350983, 61.795288], [38.089977, 55.756744], [30.262533, 59.953304], [30.41836, 59.94368], [30.029167, 59.849758], [30.319431, 60.044483], [30.614849, 59.731327], [44.082855, 56.300865], [30.463945, 59.902119], [30.460484, 59.894314], [30.46875, 59.922287], [37.728413, 55.693161], [30.231226, 59.946014], [30.24472, 59.955975], [37.396667, 55.798706], [37.758884, 55.70232], [37.758884, 55.70232], [30.227314, 59.981552], [30.216248, 59.985374], [30.243324, 59.957798], [37.92775, 59.090694], [30.24472, 59.955975], [30.306528, 59.905758], [30.354805, 59.761803], [30.306528, 59.905758], [37.917946, 59.089333], [30.434843, 59.776707], [34.357529, 61.795834], [30.354059, 60.0509], [30.323803, 59.808727], [30.399895, 59.799801], [30.564167, 59.750343], [30.319431, 60.044483], [30.398174, 59.801132], [34.357529, 61.795834], [73.449486, 61.257442], [39.179169, 51.717499], [30.402172, 59.809639], [30.304995, 59.906403], [39.781254, 54.612293], [39.159138, 51.710972], [39.781254, 54.612293], [30.110039, 59.847805], [30.396049, 60.03421], [30.067728, 59.845058], [49.330833, 53.545834], [34.344975, 61.794956], [30.319431, 60.044483], [30.305201, 59.906319], [34.350983, 61.795288], [46.001667, 51.598331], [30.256298, 59.958683], [46.001667, 51.598331], [30.446316, 59.908722], [30.277155, 59.72105], [30.169003, 59.813629], [30.332439, 59.826344], [30.111698, 59.847164], [39.873222, 59.219528], [30.494091, 59.933277], [49.77972, 53.904999], [34.357529, 61.795834], [30.398174, 59.801132], [30.491001, 59.9375], [30.444202, 59.912117], [49.346111, 53.540279], [30.431534, 59.946743], [30.111698, 59.847164], [30.169003, 59.813629], [29.822975, 59.835602], [30.398174, 59.801132], [30.321112, 59.808956], [30.614475, 59.735294], [49.77972, 53.904999], [30.24472, 59.955975], [30.162861, 59.807167], [30.247738, 59.952545], [30.333529, 60.045242], [39.838715, 57.673241], [30.320679, 60.043739], [34.354527, 61.795582], [30.323803, 59.808727], [30.222057, 59.987171], [30.304995, 59.906403], [37.917973, 59.097057], [30.279331, 59.884975], [37.905693, 59.090752], [30.304995, 59.906403], [37.905693, 59.090752], [30.222057, 59.987171], [30.243324, 59.957798], [30.222057, 59.987171], [30.266119, 59.952393], [37.912224, 59.154446], [38.087799, 55.758202], [30.495152, 59.936497], [39.846497, 59.186054], [30.157833, 59.830814], [85.136108, 52.511391], [30.215363, 59.985455], [30.256298, 59.958683], [30.247738, 59.952545], [39.838715, 57.673241], [30.311073, 60.047928], [30.416803, 60.007446], [30.301088, 59.907742], [30.304995, 59.906403], [30.155834, 59.809006], [30.448624, 59.910351], [30.30105, 59.906067], [30.222057, 59.987171], [37.722153, 44.754757], [30.304995, 59.906403], [30.614849, 59.731327], [30.304995, 59.906403], [34.360687, 61.788483], [30.678865, 59.726051], [30.155834, 59.809006], [49.354168, 53.547222], [49.472778, 53.486389], [37.728413, 55.693161], [30.306477, 59.904129], [30.491327, 59.960579], [30.168344, 59.825733], [30.301088, 59.907742], [73.453178, 61.255978], [85.177414, 52.515366], [73.456734, 61.255062], [30.448624, 59.910351], [37.858791, 55.74551], [30.111698, 59.847164], [30.388626, 59.804993], [44.943924, 53.213825], [37.923306, 59.09375], [30.222057, 59.987171], [30.496153, 59.935627], [39.79274, 54.613075], [30.245024, 59.956047], [34.391724, 61.791557], [30.311073, 60.047928], [39.832741, 54.777611], [30.230551, 59.824471], [30.443539, 59.912037], [34.379421, 61.78743], [30.336432, 59.865856], [30.321133, 59.808334], [30.231846, 59.956741], [30.16925, 59.809929], [49.354168, 53.547222], [30.222057, 59.987171], [39.846497, 59.186054], [30.321089, 59.833145], [30.10704, 59.848518], [30.319431, 60.044483], [37.920521, 59.125759], [39.179169, 51.717499], [39.175278, 51.723331], [37.773571, 55.701473], [30.155834, 59.809006], [37.917946, 59.089333], [30.188337, 59.821945], [37.917946, 59.089333], [30.444202, 59.912117], [30.294247, 59.811092], [37.986801, 55.793999], [37.917946, 59.089333], [30.244759, 59.955982], [30.32514, 59.809891], [37.65847, 44.815666], [44.890411, 53.221928], [30.155834, 59.809006], [30.301603, 59.907833], [30.163561, 59.809929], [30.181974, 59.82613], [30.398174, 59.801132], [30.309454, 59.904526], [37.759468, 55.702091], [30.446316, 59.908722], [30.460428, 59.821934], [55.119999, 51.842499], [30.433374, 60.032806], [30.434361, 59.942982], [39.179169, 51.717499], [30.241325, 59.853703], [30.301603, 59.907833], [30.306528, 59.905758], [30.603579, 59.747532], [46.003017, 51.599789], [30.456095, 59.906143], [39.185905, 51.712513], [30.311073, 60.047928], [73.45211, 61.245491], [30.24472, 59.955975], [50.803726, 61.69672], [39.76746, 54.63065], [30.113703, 59.795944], [37.75803, 55.70266], [30.485064, 59.919563], [39.788334, 54.620098], [30.155834, 59.809006], [44.952351, 53.213383], [37.917946, 59.089333], [37.762081, 55.625462], [30.244759, 59.955982], [85.194534, 52.534084], [37.917946, 59.089333], [37.917946, 59.089333], [30.093334, 59.853798], [30.155834, 59.809006], [30.433374, 60.032806], [30.434814, 60.027538], [30.319431, 60.044483], [30.319431, 60.044483], [30.305201, 59.906319], [37.423634, 55.730919], [30.465088, 59.91737], [39.846497, 59.186054], [37.917946, 59.089333], [30.330246, 59.958157], [39.943333, 57.656326], [37.905693, 59.090752], [30.238499, 59.952374], [37.721348, 55.69471], [38.084274, 55.768669], [49.354168, 53.547222], [30.434814, 60.027538], [30.456181, 60.003971], [30.10704, 59.848518], [37.646229, 55.836586], [53.218311, 56.855015], [85.17691, 52.515163], [37.646229, 55.836586], [30.10704, 59.848518], [37.905693, 59.090752], [30.246279, 59.954113], [37.768295, 55.636665], [30.434814, 60.027538], [30.495838, 59.934013], [37.968353, 55.791592], [85.17691, 52.515163], [30.304995, 59.906403], [30.497301, 59.934628], [37.917946, 59.089333], [37.602978, 55.798016], [37.905693, 59.090752], [73.454353, 61.259205], [30.304995, 59.906403], [30.074743, 59.816082], [44.08337, 56.300865], [82.921669, 55.118889], [30.315357, 59.807198], [34.344788, 61.79546], [30.41795, 60.008476], [30.321112, 59.808956], [30.304995, 59.906403], [39.943333, 57.656326], [30.093334, 59.853798], [30.252354, 59.966106], [36.568531, 50.565289], [30.162861, 59.807167], [30.28861, 59.859779], [37.884361, 59.133583], [30.495838, 59.934013], [30.155834, 59.809006], [30.301603, 59.907833], [29.943457, 59.875278], [30.252354, 59.966106], [30.326496, 59.811073], [36.575207, 50.560352], [30.398251, 59.798973], [30.304995, 59.906403], [37.905693, 59.090752], [30.497301, 59.934628], [55.119999, 51.842499], [30.325132, 60.041878], [34.357529, 61.795834], [30.43483, 59.94669], [30.434814, 60.027538], [49.360985, 53.544167], [30.426298, 59.948257], [30.990528, 59.886272], [30.495152, 59.936497], [39.233639, 51.845276], [30.678865, 59.726051], [55.956112, 54.702778], [37.912388, 59.104279], [30.262327, 59.841892], [30.436115, 59.961884], [30.300783, 59.897026], [30.432152, 59.946026], [30.495152, 59.936497], [73.411049, 61.262402], [37.917946, 59.089333], [30.319431, 60.044483], [30.446316, 59.908722], [30.20606, 59.843914], [34.357529, 61.795834], [30.330204, 60.044949], [30.077362, 59.847858], [85.203117, 52.545303], [34.359081, 61.795567], [30.495152, 59.936497], [30.614849, 59.731327], [73.457802, 61.260803], [37.917946, 59.089333], [37.911888, 59.09864], [30.24472, 59.955975], [37.917946, 59.089333], [30.237879, 59.949268], [55.126667, 51.850277], [30.446316, 59.908722], [37.639313, 55.823074], [30.446316, 59.908722], [30.16925, 59.809929], [73.399101, 61.266384], [43.921734, 56.318165], [30.162861, 59.807167], [30.43483, 59.94669], [30.311073, 60.047928], [30.163561, 59.809929], [30.494091, 59.933277], [30.215363, 59.985455], [37.636436, 55.820198], [30.433374, 60.032806], [37.636436, 55.820198], [30.24472, 59.955975], [37.636436, 55.820198], [30.4863, 59.929993], [30.319431, 60.044483], [30.431534, 59.946743], [37.636436, 55.820198], [30.222057, 59.987171], [37.917946, 59.089333], [30.614849, 59.731522], [34.398857, 61.767227], [30.444202, 59.912117], [34.316521, 61.811954], [30.067728, 59.845058], [30.24472, 59.955975], [30.111698, 59.847164], [30.321112, 59.808956], [34.357529, 61.795834], [30.458803, 59.878021], [53.2328, 56.860466], [30.093334, 59.854534], [30.222057, 59.987171], [30.432152, 59.946026], [30.308735, 59.824333], [30.222057, 59.987171], [30.304995, 59.906403], [43.911018, 56.322903], [30.222057, 59.987171], [37.917946, 59.089333], [30.111698, 59.847164], [39.838715, 57.673241], [37.965401, 55.7985], [30.398174, 59.801132], [37.333935, 55.69841], [30.434814, 60.027538], [30.093334, 59.854534], [30.678865, 59.726051], [30.222057, 59.987171], [73.460754, 61.261993], [39.943333, 57.656326], [30.289444, 59.910774], [37.785553, 55.750759], [37.858791, 55.74551], [30.420341, 59.933277], [30.309454, 59.904526], [30.222057, 59.987171], [30.581711, 59.754112], [30.678865, 59.726051], [30.093334, 59.853798], [53.328732, 56.845451], [30.678865, 59.726051], [37.858791, 55.74551], [30.304995, 59.906403], [30.222057, 59.987171], [39.838715, 57.673241], [30.10704, 59.848518], [37.951138, 59.128277], [30.614849, 59.731327], [30.304995, 59.906403], [30.093334, 59.854534], [37.917946, 59.089333], [39.828945, 59.198349], [44.082855, 56.300865], [30.282223, 59.825718], [30.394876, 59.929958], [30.215363, 59.985455], [30.103153, 59.772823], [30.256298, 59.958683], [37.917946, 59.089333], [37.883728, 55.787594], [30.444202, 59.912117], [30.244577, 59.952461], [30.222057, 59.987171], [30.43483, 59.94669], [30.412264, 59.795349], [34.350983, 61.795288], [30.226297, 59.984493], [30.434814, 60.027538], [30.304995, 59.906403], [30.370258, 60.049549], [39.87075, 59.19664], [30.162861, 59.807167], [30.24472, 59.955975], [55.119999, 51.842499], [30.16925, 59.809929], [39.838715, 57.673241], [37.728458, 55.692741], [37.721348, 55.69405], [46.001667, 51.598331], [30.497301, 59.934628], [30.323803, 59.808727], [30.678865, 59.726051], [37.910114, 59.094299], [30.314676, 59.910892], [30.321112, 59.808956], [30.24472, 59.955975], [30.311073, 60.047928], [37.9851, 55.798199], [30.20483, 59.847942], [30.309454, 59.904526], [37.858791, 55.74551], [30.222057, 59.987171], [34.436424, 61.754345], [30.269926, 59.95079], [30.155834, 59.809006], [30.125208, 59.807083], [30.155834, 59.809006], [37.6343, 55.819759], [30.304995, 59.906403], [30.569349, 59.749855], [38.073421, 55.756794], [30.216248, 59.985374], [46.001667, 51.598331], [49.472778, 53.486389], [30.250736, 59.952724], [30.222057, 59.987171], [37.726757, 55.695724], [30.326885, 59.818264], [34.357529, 61.795834], [39.788334, 54.620098], [30.321112, 59.808956], [38.042469, 55.784851], [30.326496, 59.811073], [30.426285, 59.942333], [30.427782, 60.030994], [49.472778, 53.486389], [30.372793, 59.809406], [38.087799, 55.758202], [30.222057, 59.987171], [30.24472, 59.955975], [39.913994, 57.661671], [30.309147, 59.722343], [37.9851, 55.798199], [30.313444, 60.051159], [55.971943, 54.725277], [30.215363, 59.985455], [30.434814, 60.027538], [30.163561, 59.809929], [37.9851, 55.798199], [30.155834, 59.809006], [38.054485, 55.754776], [37.905693, 59.090752], [46.001667, 51.598331], [30.434557, 60.028271], [30.491203, 59.941776], [30.398251, 59.798973], [38.081207, 55.760586], [30.30094, 59.877388], [37.449707, 55.839092], [30.319431, 60.044483], [30.321112, 59.808956], [30.269926, 59.95079], [30.228149, 59.979179], [37.917946, 59.089333], [37.86002, 55.741844], [30.416573, 60.030197], [30.509666, 59.933533], [30.398174, 59.801132], [44.08337, 56.300865], [30.215363, 59.985455], [30.678865, 59.726051], [39.79882, 54.613239], [30.244759, 59.955982], [30.323803, 59.808727], [37.720901, 55.683376], [37.917946, 59.089333], [37.905693, 59.090752], [37.334568, 55.699932], [30.24472, 59.955975], [30.24472, 59.955975], [30.301603, 59.909851], [34.357529, 61.795834], [30.444202, 59.912117], [30.433374, 60.032806], [30.320593, 60.016956], [30.154856, 59.809929], [30.418371, 59.943672], [50.806854, 61.698402], [37.79837, 44.672077], [37.982121, 55.798565], [30.222057, 59.987171], [73.456734, 61.255062], [30.093334, 59.854534], [30.320593, 60.016956], [37.917946, 59.089333], [30.324156, 60.044415], [30.495152, 59.936497], [30.304995, 59.906403], [30.162861, 59.807167], [39.883251, 57.625416], [34.291599, 61.771461], [38.088821, 55.754051], [30.24472, 59.955975], [49.361942, 53.544167], [30.245024, 59.956047], [30.245024, 59.956047], [30.434814, 60.027538], [30.491001, 59.9375], [30.310421, 59.916752], [34.344788, 61.79546], [39.943333, 57.656326], [49.475159, 53.485199], [30.24971, 59.951832], [30.321112, 59.808956], [30.433374, 60.032806], [30.398251, 59.798973], [30.678865, 59.726051], [37.905693, 59.090752], [37.661068, 55.759171], [30.434814, 60.027538], [30.162861, 59.807167], [30.318043, 60.045853], [30.316566, 60.051048], [30.39872, 59.81171], [30.168381, 59.81007], [30.614849, 59.730091], [30.222057, 59.987171], [37.928783, 59.122658], [30.43483, 59.94669], [30.614849, 59.730091], [37.858791, 55.74551], [30.256298, 59.958683], [30.315216, 59.810055], [30.230844, 59.945786], [37.759468, 55.702091], [34.357529, 61.795834], [30.446316, 59.908722], [38.093479, 55.754391], [30.448624, 59.910351], [37.917946, 59.089333], [30.497301, 59.934628], [38.994324, 55.789764], [37.89764, 59.130924], [46.001667, 51.598331], [30.319431, 60.044483], [37.985085, 55.798168], [30.093334, 59.854534], [30.316261, 60.040699], [30.409195, 59.879417], [49.354168, 53.547222], [30.304995, 59.906403], [30.323803, 59.808727], [37.9851, 55.798199], [30.123936, 59.854126], [30.432152, 59.946026], [37.923306, 59.09375], [34.355587, 61.801044], [30.495152, 59.936497], [30.10704, 59.848518], [30.10704, 59.848518], [30.432152, 59.946026], [30.43483, 59.94669], [37.444386, 55.82655], [30.272902, 60.002865], [30.31587, 60.046764], [30.240139, 59.843056], [30.422112, 59.867386], [30.093334, 59.853798], [30.111698, 59.847164], [30.429876, 59.947224], [30.10704, 59.848518], [30.464001, 59.911049], [30.111698, 59.847164], [30.34177, 59.95805], [30.33359, 60.052299], [37.917946, 59.089333], [37.905693, 59.090752], [37.770718, 55.69952], [37.761578, 55.702179], [37.965401, 55.7985], [44.0658, 56.300865], [37.985085, 55.798168], [30.111698, 59.847164], [30.155834, 59.809006], [39.179169, 51.717499], [37.917946, 59.089333], [37.654186, 55.741455], [30.43483, 59.94669], [39.79274, 54.613075], [30.168381, 59.81007], [30.168381, 59.81007], [30.313948, 60.049637], [30.319431, 60.044483], [38.073421, 55.756794], [37.917946, 59.089333], [30.222057, 59.987171], [37.923306, 59.09375], [30.155834, 59.809006], [30.093334, 59.853798], [30.093334, 59.85556], [30.319431, 60.044483], [30.092058, 59.846962], [30.156893, 59.808079], [30.231226, 59.946014], [37.856239, 55.746342], [34.354527, 61.795582], [30.448624, 59.910351], [30.434814, 60.027538], [30.283903, 59.902214], [30.495152, 59.936497], [37.917946, 59.089333], [30.429016, 60.033619], [82.947525, 55.108124], [30.446316, 59.908722], [30.331333, 60.043407], [30.155834, 59.809006], [37.636436, 55.820198], [30.495838, 59.934013], [30.28517, 59.976429], [30.440195, 59.913448], [30.315493, 59.926029], [30.266119, 59.952393], [30.111698, 59.847164], [30.319431, 60.044483], [39.788334, 54.620098], [37.334568, 55.703262], [37.334568, 55.703262], [30.300295, 59.933571], [30.107775, 59.818859], [30.297989, 59.907173], [37.642517, 55.828777], [30.378817, 60.126659], [30.222057, 59.987171], [30.321089, 59.833145], [30.321133, 59.808334], [30.328112, 59.916344], [30.32514, 59.809891], [44.082855, 56.300865], [38.000614, 59.119419], [38.000614, 59.119419], [37.917946, 59.089333], [30.325085, 59.912895], [73.456734, 61.255062], [30.245024, 59.956047], [30.301603, 59.909851], [30.427782, 60.030994], [30.215363, 59.985455], [46.001667, 51.598331], [34.354527, 61.795582], [49.354168, 53.547222], [30.093334, 59.854534], [44.039696, 56.284695], [30.378817, 60.126659], [37.917946, 59.089333], [39.826527, 59.192196], [34.344975, 61.794956], [30.398174, 59.801132], [30.162861, 59.807167], [37.883728, 55.787594], [30.446487, 59.938053], [30.222057, 59.987171], [30.34569, 59.90715], [30.235598, 59.825657], [30.321112, 59.808956], [37.917946, 59.089333], [30.316261, 60.040699], [30.315216, 59.810055], [30.315216, 59.810055], [46.003017, 51.599789], [30.294989, 59.808968], [34.344788, 61.79546], [30.277493, 59.942928], [30.323803, 59.808727], [30.319431, 60.044483], [30.163561, 59.809929], [30.50392, 59.98278], [30.24472, 59.955975], [30.432152, 59.946026], [30.572617, 59.956299], [30.434814, 60.027538], [30.093334, 59.85556], [34.357529, 61.795834], [30.325407, 59.913803], [34.357529, 61.795834], [30.332247, 60.01416], [30.516918, 59.944984], [30.319431, 60.044483], [34.357529, 61.795834], [30.50392, 59.98278], [30.398174, 59.801132], [37.910114, 59.094299], [30.111698, 59.847164], [30.222057, 59.987171], [30.446819, 59.886539], [30.296774, 59.859055], [30.093334, 59.854534], [86.75576, 53.86282], [30.434338, 59.953945], [44.048553, 56.288578], [30.43158, 59.946812], [30.022676, 59.705574], [30.495152, 59.936497], [30.40671, 59.803234], [73.457802, 61.260803], [37.728863, 44.737473], [30.30094, 59.877388], [46.001667, 51.598331], [73.457413, 61.313774], [30.213709, 60.142811], [39.95755, 57.658764], [39.90575, 59.200207], [30.576466, 59.741688], [30.435919, 59.945473], [37.930721, 59.103138], [55.996666, 54.696945], [44.082855, 56.300865], [85.183609, 52.52639], [38.073532, 55.761044], [34.357529, 61.795834], [37.883728, 55.787594], [30.50952, 59.977268], [49.361942, 53.544167], [30.177652, 59.819408], [30.448624, 59.910351], [30.222057, 59.987171], [37.9851, 55.798199], [30.163561, 59.809929], [30.303675, 59.893383], [30.233976, 59.995193], [30.323803, 59.808727], [73.457802, 61.260803], [34.357529, 61.795834], [30.448347, 60.030094], [30.321133, 59.808334], [30.464399, 60.047668], [30.464399, 60.047668], [30.08135, 59.740864], [34.395329, 61.784279], [37.905693, 59.090752], [30.428268, 59.975586], [30.227314, 59.981552], [30.134542, 59.834881], [37.917946, 59.089333], [30.269926, 59.95079], [30.155834, 59.809006], [30.403225, 59.804989], [30.1635, 59.824326], [30.448624, 59.910351], [30.155834, 59.809006], [30.448624, 59.910351], [30.215363, 59.985455], [30.434814, 60.027538], [30.252354, 59.966106], [30.403225, 59.804989], [30.798817, 60.316898], [34.357529, 61.795834], [34.354527, 61.795582], [30.217274, 59.952969], [39.717949, 54.650875], [39.79274, 54.613075], [37.982121, 55.798565], [30.3141, 60.047661], [30.376097, 59.804626], [49.328609, 53.528332], [73.458878, 61.261238], [30.301603, 59.909851], [30.444202, 59.912117], [30.448624, 59.910351], [44.943333, 53.223331], [37.923306, 59.09375], [30.327559, 59.862499], [37.968353, 55.791592], [30.24472, 59.955975], [30.315216, 59.810055], [49.356201, 53.543888], [82.921669, 55.118889], [30.315357, 59.807198], [37.905693, 59.090752], [30.321112, 59.808956], [30.416573, 60.030197]]"

type Mock struct {
	data []float64
}

type jsonPoints [][]float64

func NewMock() *Mock {
	return &Mock{
		data: make([]float64, 0),
	}
}

func (o *Mock) GetData(query string) ([][]float64, error) {

	var points jsonPoints

	var err = json.Unmarshal([]byte(jsonPointsData), &points)

	return points, err
}
