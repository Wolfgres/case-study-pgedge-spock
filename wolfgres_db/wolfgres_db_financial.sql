-- Create Schema

ALTER SCHEMA wfg OWNER TO wolfgres_user;

CREATE TABLE wfg.account (
    account_id integer NOT NULL,
    customer_id integer,
    account_type_id integer,
    balace double precision
);


ALTER TABLE wfg.account OWNER TO wolfgres_user;


CREATE SEQUENCE wfg.account_account_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE wfg.account_account_id_seq OWNER TO wolfgres_user;


ALTER SEQUENCE wfg.account_account_id_seq OWNED BY wfg.account.account_id;



CREATE TABLE wfg.customer (
    customer_id integer NOT NULL,
    name character varying(255) NOT NULL,
    address character varying(255) NOT NULL,
    website character varying(255) NOT NULL,
    credit_limit double precision NOT NULL
);


ALTER TABLE wfg.customer OWNER TO wolfgres_user;


CREATE TABLE wfg.operation (
    operation_id integer NOT NULL,
    name character varying(50),
    description text
);


ALTER TABLE wfg.operation OWNER TO wolfgres_user;


CREATE SEQUENCE wfg.operation_operation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE wfg.operation_operation_id_seq OWNER TO wolfgres_user;


ALTER SEQUENCE wfg.operation_operation_id_seq OWNED BY wfg.operation.operation_id;



CREATE TABLE wfg.transaction (
    transaction_id integer NOT NULL,
    account_id integer,
    operation_id integer,
    mount double precision,
    date timestamp without time zone
);


ALTER TABLE wfg.transaction OWNER TO wolfgres_user;


CREATE SEQUENCE wfg.transaction_transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE wfg.transaction_transaction_id_seq OWNER TO wolfgres_user;

ALTER SEQUENCE wfg.transaction_transaction_id_seq OWNED BY wfg.transaction.transaction_id;



ALTER TABLE ONLY wfg.account ALTER COLUMN account_id SET DEFAULT nextval('wfg.account_account_id_seq'::regclass);



ALTER TABLE ONLY wfg.operation ALTER COLUMN operation_id SET DEFAULT nextval('wfg.operation_operation_id_seq'::regclass);



ALTER TABLE ONLY wfg.transaction ALTER COLUMN transaction_id SET DEFAULT nextval('wfg.transaction_transaction_id_seq'::regclass);



COPY wfg.account (account_id, customer_id, account_type_id, balace) FROM stdin;
\.



COPY wfg.customer (customer_id, name, address, website, credit_limit) FROM stdin;
1	Tazz	Maple Wood Park 28	tazz.global	75161.4282
2	Topdrive	Kenwood Trail 13	topdrive.info	91480.2339
3	Yakijo	Cottonwood Street 10	yakijo.biz	625851.1481
4	Browsetype	Browning Parkway 83	browsetype.site	914489.6477
5	Gabtype	Mosinee Drive 26	gabtype.support	207078.4459
6	Photobug	Aberg Pass 93	photobug.solutions	383007.8109
7	Rhynoodle	Dunning Court 91	rhynoodle.support	657781.417
8	Skidoo	Park Meadow Drive 3	skidoo.org	302429.3872
9	Zoomdog	Northfield Drive 99	zoomdog.biz	117979.0628
10	Linklinks	Declaration Crossing 62	linklinks.biz	245552.1986
11	Gabtune	Sunfield Street 99	gabtune.app	809597.2417
12	Leenti	Sherman Park 20	leenti.global	304653.429
13	Katz	Maryland Junction 42	katz.global	243237.3976
14	Mycat	Saint Paul Lane 81	mycat.xyz	36859.9406
15	Bluezoom	Sutherland Trail 12	bluezoom.global	833904.4564
16	Topdrive	Fuller Circle 92	topdrive.site	771903.7015
17	Ooba	Pearson Circle 88	ooba.me	260631.8087
18	Jayo	Kennedy Court 10	jayo.global	2894.6094
19	Vidoo	Crownhardt Crossing 80	vidoo.xyz	525691.0929
20	Izio	Dixon Court 73	izio.online	47191.9153
21	Meejo	Texas Lane 26	meejo.net	706872.4062
22	Camimbo	Raven Junction 29	camimbo.solutions	361908.5547
23	LiveZ	Holy Cross Court 97	livez.io	770270.1435
24	Gigabox	Doe Crossing Circle 88	gigabox.xyz	564377.287
25	Bluejam	Larry Lane 43	bluejam.co	565235.2394
26	Twitternation	Haas Drive 3	twitternation.me	755682.4315
27	Skippad	Myrtle Terrace 46	skippad.site	66123.4238
28	Shufflebeat	Derek Lane 2	shufflebeat.me	160089.7771
29	Oyonder	Mandrake Hill 11	oyonder.biz	224872.7999
30	Wikivu	Sachs Point 81	wikivu.support	57878.0629
31	Mymm	Huxley Circle 45	mymm.xyz	481484.9678
32	Vinder	Farwell Court 35	vinder.support	492894.1295
33	Kare	Graedel Court 10	kare.online	943438.1552
34	Tanoodle	Springs Junction 32	tanoodle.support	846859.2877
35	Fanoodle	Kinsman Circle 73	fanoodle.net	352332.925
36	Linkbuzz	Spohn Road 10	linkbuzz.app	599998.9219
37	Gabtype	Spaight Street 42	gabtype.org	617633.5585
38	Fliptune	Waywood Lane 36	fliptune.support	383508.1552
39	Abata	Hovde Junction 13	abata.biz	932276.0311
40	Jabberstorm	Menomonie Trail 13	jabberstorm.support	725034.1797
41	Roombo	Kenwood Pass 22	roombo.solutions	294882.391
42	Eamia	Stone Corner Point 62	eamia.site	898699.9451
43	Rhyloo	East Road 93	rhyloo.site	64415.6816
44	Flipstorm	Armistice Lane 92	flipstorm.me	193057.3915
45	Flipopia	Karstens Center 97	flipopia.io	808282.0938
46	Jabbersphere	Tomscot Way 72	jabbersphere.app	958346.1544
47	Yombu	Russell Place 43	yombu.com	571353.6642
48	Realcube	Welch Lane 22	realcube.org	878349.761
49	Lazzy	Daystar Pass 75	lazzy.io	534216.528
50	Feedbug	Lawn Parkway 16	feedbug.co	493584.9331
51	Bluezoom	Melrose Junction 8	bluezoom.support	509999.8037
52	Camido	Northwestern Lane 78	camido.com	692874.3964
53	Voolith	New Castle Plaza 77	voolith.com	201011.3951
54	Topicblab	Prentice Road 85	topicblab.info	200989.5316
55	Skajo	Clarendon Center 96	skajo.net	784693.9371
56	Wordtune	Northfield Alley 51	wordtune.online	451517.742
57	Voomm	Porter Terrace 67	voomm.org	153457.7273
58	Dablist	Towne Road 42	dablist.net	648976.8488
59	Einti	Colorado Plaza 4	einti.io	996370.1705
60	Pixonyx	Comanche Circle 70	pixonyx.org	422876.7421
61	Oyoloo	Jenna Point 0	oyoloo.io	292409.1287
62	Skaboo	Hovde Center 95	skaboo.online	623686.2538
63	Eimbee	Fordem Street 73	eimbee.co	754456.734
64	Twimbo	Boyd Pass 10	twimbo.xyz	926809.5769
65	Kaymbo	2nd Junction 86	kaymbo.online	262761.8753
66	Devshare	Weeping Birch Drive 94	devshare.me	758960.9369
67	Realmix	Meadow Ridge Way 84	realmix.xyz	155572.2785
68	Livetube	Hallows Pass 56	livetube.site	673071.7883
69	Avamm	Magdeline Court 70	avamm.global	486097.8566
70	Flipbug	Sommers Lane 21	flipbug.net	166597.9953
71	Oyoba	Duke Crossing 45	oyoba.biz	177912.0145
72	Zoombox	Spenser Alley 14	zoombox.io	805702.0206
73	Ntag	Dayton Terrace 30	ntag.io	748622.9993
74	Twiyo	Cottonwood Pass 96	twiyo.app	110163.6391
75	Photolist	Portage Parkway 74	photolist.io	137226.0315
76	Ooba	Carey Pass 25	ooba.online	748390.267
77	Ailane	Doe Crossing Plaza 74	ailane.com	671388.0817
78	Eazzy	Grayhawk Trail 70	eazzy.online	62858.4158
79	Realmix	Kenwood Crossing 38	realmix.xyz	662217.6376
80	JumpXS	Forest Pass 76	jumpxs.solutions	35635.2167
81	Linkbuzz	Arrowood Avenue 24	linkbuzz.com	815700.8603
82	Tagchat	Southridge Lane 20	tagchat.info	713457.0183
83	Cogilith	Corry Junction 29	cogilith.org	899857.0416
84	Kwilith	Browning Way 29	kwilith.me	965574.4198
85	Riffpedia	Roxbury Terrace 53	riffpedia.online	326292.0283
86	Lazzy	Pankratz Place 49	lazzy.info	882295.3903
87	Browsecat	Thackeray Circle 28	browsecat.support	408149.8052
88	Aibox	Emmet Pass 59	aibox.xyz	739861.4133
89	Realbridge	Blaine Terrace 89	realbridge.support	969848.8643
90	Omba	Merchant Plaza 86	omba.app	946286.9123
91	Viva	Bluejay Center 68	viva.online	417523.8863
92	Brightdog	Village Green Terrace 22	brightdog.solutions	43247.9614
93	Voonder	Montana Plaza 99	voonder.net	623918.664
94	Gigabox	Buhler Drive 52	gigabox.biz	116167.0683
95	Myworks	Reinke Drive 25	myworks.global	740472.9151
96	Kare	Twin Pines Street 30	kare.info	325714.032
97	Wikivu	Meadow Valley Circle 21	wikivu.net	516036.7636
98	Gevee	Schurz Lane 72	gevee.xyz	382974.8896
99	Realfire	Lien Crossing 91	realfire.com	726116.6874
100	JumpXS	Morning Hill 8	jumpxs.me	158957.2617
101	Zoonder	Oak Valley Lane 82	zoonder.online	560089.3141
102	Mycat	Pennsylvania Junction 86	mycat.org	203443.1333
103	Yadel	Briar Crest Hill 83	yadel.co	280628.2596
104	Flipstorm	Forest Drive 85	flipstorm.site	967795.1895
105	Camido	Summerview Street 22	camido.xyz	269995.9406
106	Katz	Moland Plaza 29	katz.me	636899.1738
107	Babbleset	Bashford Trail 15	babbleset.site	871352.0071
108	Pixonyx	Kipling Lane 33	pixonyx.com	121443.1414
109	Divavu	Lukken Park 92	divavu.com	306933.8276
110	Tanoodle	Paget Junction 25	tanoodle.xyz	778326.8018
111	Tekfly	Morrow Place 76	tekfly.xyz	561698.0156
112	Eazzy	Meadow Ridge Park 95	eazzy.net	903460.0785
113	Skipstorm	Caliangt Crossing 83	skipstorm.xyz	776427.4173
114	Twitterbeat	Petterle Road 18	twitterbeat.online	621252.7448
115	Fiveclub	Brown Lane 42	fiveclub.app	959225.8638
116	Twitterlist	Lighthouse Bay Point 66	twitterlist.biz	512629.3712
117	Aimbu	Bellgrove Trail 61	aimbu.org	374657.5065
118	Feedfire	Gale Way 16	feedfire.site	812752.875
119	Skajo	Mosinee Park 46	skajo.org	145105.2241
120	Oyope	Northridge Terrace 44	oyope.xyz	228080.5964
121	Eadel	Shopko Center 69	eadel.app	433725.7367
122	Flashset	New Castle Court 43	flashset.xyz	838626.8155
123	Zoombeat	Shopko Way 10	zoombeat.global	241375.7867
124	Buzzshare	Kings Pass 56	buzzshare.com	125234.5232
125	Lazzy	Jenifer Point 49	lazzy.net	347058.0863
126	Abatz	Forest Dale Circle 75	abatz.app	574538.4906
127	Realcube	Northfield Point 35	realcube.support	75537.7496
128	Aivee	Crest Line Lane 6	aivee.org	60198.3197
129	Tagchat	Jana Junction 79	tagchat.site	411522.1384
130	Thoughtmix	Dryden Hill 6	thoughtmix.online	895506.6834
131	Demimbu	Magdeline Court 48	demimbu.online	216172.62
132	Skajo	Dovetail Way 6	skajo.org	361383.3188
133	Geba	Huxley Street 8	geba.app	789967.2023
134	Tazz	Summerview Junction 72	tazz.site	50063.9172
135	Bubbletube	Clove Point 65	bubbletube.info	64076.16
136	Thoughtworks	Pond Center 19	thoughtworks.net	179155.8022
137	Twitterbeat	Bowman Street 30	twitterbeat.me	126332.2245
138	Thoughtmix	Morningstar Park 30	thoughtmix.co	985333.6177
139	Lajo	Bowman Street 19	lajo.co	574747.4585
140	Gabtype	Westridge Lane 41	gabtype.me	100653.8421
141	Blogspan	Rutledge Hill 51	blogspan.global	104406.77
142	Youspan	Beilfuss Junction 50	youspan.io	459202.2578
143	Yabox	Myrtle Way 14	yabox.me	386166.3863
144	Flashset	Evergreen Road 83	flashset.org	936272.0197
145	Riffwire	Shoshone Pass 29	riffwire.site	796829.2894
146	Feedmix	Packers Plaza 84	feedmix.co	748728.2564
147	Leexo	Ryan Park 64	leexo.site	473805.1443
148	Jayo	Summit Circle 65	jayo.xyz	846493.2591
149	Ozu	Loeprich Crossing 11	ozu.global	731316.2184
150	Tagcat	Tennyson Crossing 7	tagcat.support	670004.5924
151	Voolia	Dryden Road 50	voolia.io	308688.3256
152	Oloo	Hagan Circle 94	oloo.online	157795.4089
153	Kare	Blaine Junction 23	kare.xyz	919526.4999
154	Demimbu	Sunfield Alley 56	demimbu.me	692145.8352
155	Roombo	Luster Plaza 99	roombo.org	386493.2466
156	Meevee	Oak Valley Street 25	meevee.xyz	140750.9067
157	Jabbercube	Scoville Junction 14	jabbercube.app	317618.9863
158	Avamm	Schmedeman Crossing 86	avamm.global	775736.537
159	Mynte	Crest Line Circle 3	mynte.org	577400.089
160	Dynazzy	Havey Parkway 75	dynazzy.xyz	555105.5217
161	Ntag	Basil Junction 76	ntag.net	770609.2722
162	Youfeed	Dovetail Junction 19	youfeed.solutions	293799.8206
163	Realfire	Hauk Terrace 90	realfire.com	488760.4654
164	Jatri	Arrowood Parkway 20	jatri.support	844129.7581
165	Browseblab	Swallow Avenue 38	browseblab.site	786263.0403
166	Centizu	Lunder Hill 40	centizu.co	234980.3896
167	DabZ	Spohn Circle 3	dabz.me	624414.4454
168	Minyx	Dexter Plaza 97	minyx.io	779326.944
169	Jabberbean	Sycamore Crossing 49	jabberbean.org	833677.3912
170	Zoomlounge	North Crossing 96	zoomlounge.me	991941.9813
171	Thoughtblab	Old Shore Road 7	thoughtblab.org	396064.0648
172	Flashpoint	Brickson Park Drive 56	flashpoint.xyz	868299.8233
173	Blogpad	Walton Circle 21	blogpad.site	326500.0568
174	Quire	Ryan Alley 89	quire.global	665162.1724
175	Realbridge	Pond Junction 9	realbridge.net	302260.1766
176	Gigaclub	Lyons Park 10	gigaclub.co	404643.0108
177	Fanoodle	Annamark Drive 9	fanoodle.com	738564.9473
178	Quire	Nevada Point 50	quire.biz	431622.9795
179	Eamia	Rieder Parkway 92	eamia.site	888270.5735
180	Flashset	Mandrake Alley 21	flashset.online	48131.1067
181	Twiyo	Parkside Place 6	twiyo.io	381120.0734
182	BlogXS	Sutteridge Junction 41	blogxs.site	589379.3789
183	Yadel	Bartelt Avenue 80	yadel.net	393423.8941
184	Tagchat	Heffernan Crossing 3	tagchat.me	935276.7558
185	Ainyx	Norway Maple Place 65	ainyx.global	864149.7754
186	Wikibox	Nevada Place 25	wikibox.com	184010.5224
187	Photobug	Magdeline Pass 42	photobug.site	855165.5126
188	Avamba	Lakeland Park 92	avamba.net	207348.3552
189	Meeveo	Bluejay Avenue 40	meeveo.io	404446.8853
190	Meemm	Dahle Terrace 96	meemm.app	697509.7616
191	Demizz	Miller Drive 23	demizz.io	309423.0856
192	Devcast	Judy Alley 28	devcast.global	951441.177
193	Wordware	Meadow Valley Road 66	wordware.app	754839.3001
194	Fliptune	Novick Drive 60	fliptune.me	993415.9846
195	Wordware	Larry Crossing 3	wordware.support	476291.44
196	Babblestorm	Kenwood Trail 42	babblestorm.co	448722.3033
197	Bubblemix	Independence Place 46	bubblemix.online	950264.6463
198	Oyoloo	Larry Center 27	oyoloo.online	571758.4824
199	Aimbu	Crescent Oaks Drive 54	aimbu.global	800671.6037
200	Quinu	Pierstorff Avenue 84	quinu.com	510049.4769
201	Wikivu	Clarendon Court 18	wikivu.app	751710.3063
202	Yabox	Crowley Circle 51	yabox.org	779520.951
203	Rhyzio	Brentwood Lane 99	rhyzio.net	773902.0343
204	Bubbletube	Hagan Circle 69	bubbletube.io	153080.2734
205	Bluejam	Talmadge Drive 38	bluejam.biz	725129.5821
206	Aimbu	Di Loreto Circle 44	aimbu.org	419914.0996
207	Riffpath	Wayridge Center 0	riffpath.biz	625238.1887
208	Livetube	Reinke Road 88	livetube.site	405971.2812
209	Eayo	Walton Court 0	eayo.com	984873.7581
210	Wikivu	Dwight Avenue 40	wikivu.online	897219.7142
211	Innotype	Mcbride Street 31	innotype.org	83326.5735
212	Nlounge	Sachtjen Terrace 0	nlounge.co	496523.928
213	Voomm	Butternut Street 9	voomm.global	353892.7524
214	Jazzy	Thierer Road 8	jazzy.global	240012.842
215	Zooxo	4th Center 30	zooxo.info	767171.742
216	Tagpad	Thompson Park 54	tagpad.xyz	704280.1642
217	Skinix	David Center 8	skinix.co	73728.3704
218	Gabtune	Hovde Junction 86	gabtune.org	55473.0467
219	Jabberbean	Green Ridge Lane 58	jabberbean.io	9173.613
220	Ainyx	Grasskamp Avenue 34	ainyx.solutions	315362.7138
221	Jaxbean	Jenifer Parkway 2	jaxbean.global	678108.0242
222	Zava	Moland Crossing 41	zava.support	424741.5336
223	Ozu	Thierer Road 7	ozu.solutions	728343.5958
224	Buzzdog	Debra Hill 41	buzzdog.solutions	938188.0329
225	Fivespan	Brown Junction 27	fivespan.biz	969274.3985
226	Oyoba	Gina Street 5	oyoba.me	239712.8144
227	Youspan	Carioca Place 7	youspan.online	895056.7042
228	Yakitri	Leroy Way 98	yakitri.global	374588.8323
229	Bubblebox	Arkansas Drive 34	bubblebox.io	267359.1937
230	Plajo	Iowa Park 69	plajo.org	731046.7687
231	Voonte	Lukken Center 5	voonte.app	996049.1843
232	InnoZ	Hallows Drive 90	innoz.xyz	533881.7468
233	Youspan	Sycamore Road 70	youspan.app	534609.0091
234	Brightdog	Dapin Alley 72	brightdog.net	919935.2497
235	InnoZ	Meadow Ridge Alley 64	innoz.global	925716.8367
236	Blognation	Parkside Court 85	blognation.support	343858.0079
237	Zoonder	Lien Place 88	zoonder.app	838338.1627
238	Tekfly	Arkansas Plaza 88	tekfly.org	140147.1921
239	Jaxworks	Londonderry Way 45	jaxworks.net	545295.2824
240	Meevee	Weeping Birch Park 5	meevee.com	298860.009
241	Dynabox	Evergreen Trail 62	dynabox.site	795016.7566
242	Gigabox	Bartelt Trail 54	gigabox.xyz	305952.205
243	Meevee	Fremont Plaza 26	meevee.com	215986.3989
244	Feedfish	Sullivan Park 43	feedfish.io	153784.0353
245	Brightbean	Westend Plaza 66	brightbean.app	376489.2752
246	Linktype	Superior Pass 2	linktype.net	306956.3099
247	Oyoyo	Myrtle Pass 89	oyoyo.solutions	504509.4009
248	Livepath	Sundown Junction 23	livepath.global	837150.7078
249	Yodel	Burning Wood Place 10	yodel.xyz	242389.9799
250	Topicshots	Pawling Junction 44	topicshots.co	502730.6451
251	Rhybox	Bultman Center 6	rhybox.global	279924.3668
252	Browsecat	Iowa Park 63	browsecat.com	856741.1582
253	Oyope	Brentwood Drive 2	oyope.org	920958.439
254	Abata	Ridge Oak Drive 50	abata.xyz	226800.9657
255	JumpXS	Maple Hill 23	jumpxs.biz	875712.6123
256	Oyoloo	Mcguire Center 96	oyoloo.online	675409.69
257	Bubblebox	Victoria Parkway 63	bubblebox.com	394523.8236
258	Meejo	Barby Way 50	meejo.app	28892.6959
259	Janyx	Waubesa Park 68	janyx.global	847369.9259
260	Thoughtsphere	Oneill Lane 1	thoughtsphere.org	157621.1665
261	Ntag	Fieldstone Circle 82	ntag.site	711614.4811
262	Voolia	Messerschmidt Street 16	voolia.support	581780.0654
263	Linkbridge	Rigney Pass 71	linkbridge.biz	412171.7214
264	Avamba	Lillian Lane 71	avamba.com	288349.8483
265	Npath	Harper Park 41	npath.solutions	711028.4109
266	BlogXS	Norway Maple Parkway 23	blogxs.solutions	675956.0146
267	Leexo	Logan Court 5	leexo.info	199497.6058
268	Jatri	Dakota Avenue 93	jatri.io	938791.5322
269	Dabtype	Harbort Road 86	dabtype.global	786267.3059
270	Flipstorm	Hoepker Alley 71	flipstorm.co	569904.8161
271	Zoomcast	Anhalt Street 4	zoomcast.solutions	38108.8748
272	Devcast	Kropf Terrace 84	devcast.online	350577.5943
273	Zoozzy	Portage Way 11	zoozzy.net	3720.9603
274	Zooxo	Dunning Center 9	zooxo.net	923341.0911
275	Devcast	Washington Alley 58	devcast.net	251831.9855
276	Realcube	Messerschmidt Crossing 98	realcube.site	603167.1641
277	Buzzshare	Melby Junction 98	buzzshare.com	670740.6145
278	Youspan	Oakridge Street 35	youspan.xyz	521477.0663
279	Wordpedia	Sugar Crossing 41	wordpedia.biz	206999.6002
280	Reallinks	Transport Park 89	reallinks.net	669026.9828
281	Skyvu	Stuart Way 83	skyvu.org	536490.6701
282	Brainlounge	Glendale Pass 30	brainlounge.org	978997.987
283	Tagchat	Derek Plaza 36	tagchat.org	654800.0261
284	Fivebridge	Johnson Road 92	fivebridge.info	71385.5331
285	Zooxo	Park Meadow Road 98	zooxo.com	314927.9317
286	Reallinks	Graedel Terrace 42	reallinks.site	156325.7623
287	Dabjam	Oriole Way 91	dabjam.online	69551.8478
288	Rooxo	Pennsylvania Court 36	rooxo.support	842069.2825
289	Linkbridge	Oak Plaza 23	linkbridge.biz	313045.1997
290	Talane	Pierstorff Parkway 31	talane.org	750779.4055
291	Skibox	Longview Center 64	skibox.online	317481.576
292	Photobug	Golf Course Center 31	photobug.net	734659.8892
293	Brainsphere	Buell Crossing 96	brainsphere.me	580606.3489
294	Latz	Clarendon Street 65	latz.org	557502.5696
295	Gabcube	Saint Paul Alley 97	gabcube.io	688191.9016
296	Voomm	Steensland Pass 29	voomm.biz	719521.0491
297	Avamm	Glacier Hill Center 49	avamm.info	252033.5167
298	Dabvine	Carioca Point 26	dabvine.io	622742.5835
299	Twitterlist	Burrows Hill 87	twitterlist.co	863202.3865
300	Zoombox	Kings Alley 23	zoombox.support	607754.4755
301	Minyx	Amoth Plaza 4	minyx.solutions	90569.7239
302	Twitterlist	Shoshone Pass 55	twitterlist.io	415668.2778
303	Dabfeed	Fisk Junction 67	dabfeed.biz	676736.0278
304	Twimbo	Manitowish Way 82	twimbo.app	48389.4023
305	Mudo	Lawn Street 18	mudo.com	886840.4125
306	Blogtags	Monica Circle 4	blogtags.app	882403.6305
307	Yombu	Starling Avenue 39	yombu.app	780529.6189
308	Jabberstorm	Anzinger Place 54	jabberstorm.org	934897.0692
309	Kazu	Jay Terrace 10	kazu.online	441768.3689
310	Feedbug	Debs Trail 99	feedbug.org	606170.5076
311	Brightbean	Shoshone Park 87	brightbean.online	793917.7567
312	Thoughtmix	Algoma Road 26	thoughtmix.me	540462.2803
313	Photolist	School Way 60	photolist.app	513989.4048
314	Fanoodle	Brown Trail 63	fanoodle.net	644477.8907
315	Camido	Oakridge Terrace 39	camido.site	51884.6703
316	Feednation	7th Parkway 21	feednation.support	926152.6005
317	Topiclounge	Vahlen Terrace 58	topiclounge.me	144133.557
318	Twiyo	Golf Course Road 18	twiyo.io	175420.2539
319	Vinte	Mallard Lane 23	vinte.me	293132.4392
320	Topdrive	Manufacturers Trail 9	topdrive.xyz	756091.8629
321	Podcat	Twin Pines Parkway 71	podcat.global	401500.292
322	Zoombox	Bowman Point 32	zoombox.online	454399.6847
323	Topicshots	Bluejay Avenue 89	topicshots.xyz	815809.945
324	Flashset	Nelson Alley 77	flashset.io	782107.458
325	Gevee	Johnson Way 27	gevee.org	367948.3428
326	Quinu	Schiller Way 68	quinu.co	353840.2693
327	Kazio	Macpherson Parkway 25	kazio.xyz	247046.1008
328	Mudo	Fieldstone Center 4	mudo.org	372252.6516
329	Trilia	Tomscot Park 15	trilia.biz	394940.6559
330	Browsecat	Meadow Vale Drive 74	browsecat.solutions	787397.0748
331	Edgewire	Moulton Junction 14	edgewire.co	208516.5188
332	Innojam	Eastlawn Drive 84	innojam.co	777109.6116
333	Ailane	Dixon Lane 4	ailane.solutions	252333.1893
334	Mymm	Ridge Oak Junction 31	mymm.xyz	199440.7645
335	Omba	Hintze Circle 92	omba.biz	655179.0102
336	JumpXS	Becker Court 58	jumpxs.org	560596.9689
337	Teklist	Coleman Junction 68	teklist.me	139170.3837
338	Dabvine	Anzinger Drive 86	dabvine.io	541316.9572
339	Pixonyx	Lake View Center 76	pixonyx.io	648813.4987
340	Voolith	Reinke Pass 47	voolith.solutions	426726.5443
341	Fiveclub	Linden Parkway 22	fiveclub.org	306829.4374
342	Skalith	Artisan Road 31	skalith.global	991770.1115
343	Meemm	Summerview Alley 13	meemm.support	162780.2016
344	Gigashots	Huxley Lane 30	gigashots.xyz	847999.4342
345	Livetube	Anzinger Crossing 89	livetube.global	232914.6636
346	Innojam	Granby Center 73	innojam.online	789965.4111
347	Eayo	Badeau Terrace 6	eayo.com	840625.7449
348	Quimm	Merchant Drive 28	quimm.support	703233.8428
349	Skidoo	Anhalt Court 16	skidoo.solutions	68706.8176
350	Feedmix	Heffernan Trail 83	feedmix.online	30188.4647
351	Zazio	Menomonie Pass 89	zazio.co	30137.1994
352	Kwinu	Monica Street 35	kwinu.info	37972.5587
353	Blogpad	Dennis Crossing 56	blogpad.co	58157.5431
354	Zoovu	Namekagon Pass 37	zoovu.io	689554.8521
355	Eadel	Ramsey Alley 21	eadel.site	648044.8444
356	Pixoboo	Bowman Avenue 93	pixoboo.biz	390295.6656
357	Skivee	Fordem Street 30	skivee.global	540368.1047
358	Edgepulse	Oak Crossing 47	edgepulse.info	856218.3103
359	Jayo	Delladonna Hill 25	jayo.info	540118.1256
360	Nlounge	Sage Point 99	nlounge.app	309871.5331
361	Dabvine	Center Plaza 50	dabvine.site	315066.5024
362	Cogilith	Buena Vista Avenue 62	cogilith.io	775218.4119
363	BlogXS	Green Ridge Road 53	blogxs.co	578232.6302
364	Brightdog	Hanson Lane 41	brightdog.net	766240.2167
365	Photofeed	Mendota Terrace 3	photofeed.biz	128141.8563
366	Twitterlist	Calypso Plaza 42	twitterlist.global	402247.1623
367	Thoughtstorm	Leroy Point 7	thoughtstorm.online	266920.6701
368	Reallinks	Bunting Trail 8	reallinks.support	574144.2287
369	Kwideo	Texas Street 67	kwideo.info	626788.9368
370	Voonte	Jay Hill 43	voonte.support	214079.3191
371	Rooxo	Shelley Point 6	rooxo.xyz	890242.3453
372	Tavu	Luster Drive 78	tavu.support	865616.1362
373	Thoughtsphere	Loftsgordon Road 23	thoughtsphere.me	137164.4181
374	Voonte	Mallory Trail 69	voonte.info	771270.3891
375	Latz	Transport Way 87	latz.solutions	182135.9666
376	Kanoodle	Katie Parkway 13	kanoodle.me	326672.0918
377	Voonte	Clove Parkway 30	voonte.biz	278365.5049
378	Yoveo	Forest Run Court 57	yoveo.xyz	767349.0811
379	Livetube	Merry Place 63	livetube.site	290665.6127
380	Gigaclub	Loomis Park 85	gigaclub.com	283959.1558
381	Eidel	Bobwhite Crossing 65	eidel.global	594939.3891
382	Shufflebeat	Fremont Drive 35	shufflebeat.xyz	235770.6306
383	Latz	Lakewood Plaza 56	latz.online	502612.4635
384	Flipbug	Gerald Court 13	flipbug.net	202397.5055
385	Devpulse	Riverside Crossing 14	devpulse.io	525769.197
386	Gabtype	Truax Center 51	gabtype.co	415267.8022
387	Meevee	Warbler Road 49	meevee.global	641941.5613
388	Brainverse	Arkansas Plaza 90	brainverse.site	760241.4196
389	Livefish	Tony Hill 49	livefish.solutions	456815.6169
390	Cogidoo	Golf View Place 97	cogidoo.global	777181.3913
391	Voomm	Forest Street 71	voomm.xyz	283093.8856
392	Cogilith	Calypso Parkway 16	cogilith.app	70154.3387
393	Avamba	Nelson Way 12	avamba.xyz	620069.2083
394	Dynava	Kim Hill 92	dynava.app	830950.3727
395	Mynte	Kings Trail 13	mynte.support	862389.795
396	Katz	Mockingbird Way 9	katz.com	829601.2617
397	Gigaclub	John Wall Plaza 63	gigaclub.com	914515.5677
398	Browsezoom	School Alley 59	browsezoom.online	863626.8648
399	Realcube	Gerald Place 53	realcube.me	798449.2513
400	Edgeify	Debs Avenue 91	edgeify.co	831972.0957
401	Zazio	Spohn Hill 30	zazio.org	507014.2281
402	Realpoint	Kingsford Way 23	realpoint.biz	61764.2455
403	Skinder	Vernon Parkway 21	skinder.io	998450.5994
404	Jaxbean	Arizona Parkway 45	jaxbean.com	99504.6585
405	Wordify	Stang Alley 85	wordify.io	53402.9106
406	Feedmix	Drewry Alley 57	feedmix.solutions	762723.5385
407	Oyope	Rutledge Terrace 58	oyope.xyz	513293.937
408	Brainbox	Spaight Park 48	brainbox.com	915378.0896
409	Rhyloo	Mosinee Avenue 9	rhyloo.io	379335.988
410	Voonyx	Butternut Terrace 54	voonyx.io	63876.8467
411	Buzzshare	Roth Road 27	buzzshare.xyz	155954.9744
412	Aimbu	Butternut Court 8	aimbu.solutions	497763.7124
413	Twinte	Fulton Park 51	twinte.me	941544.6274
414	Zoombeat	Comanche Alley 58	zoombeat.solutions	663562.4009
415	Tazzy	Kipling Junction 53	tazzy.net	852692.149
416	Dazzlesphere	Chive Point 40	dazzlesphere.me	209663.9472
417	Quaxo	Chive Park 59	quaxo.xyz	938561.6824
418	Rhynyx	Thackeray Crossing 59	rhynyx.org	52945.8691
419	Photospace	Oak Way 7	photospace.solutions	162663.6786
420	Jabbercube	Menomonie Hill 37	jabbercube.xyz	97031.3332
421	Oloo	Village Road 46	oloo.co	48387.4997
422	Dynava	Londonderry Point 99	dynava.online	636878.611
423	Devcast	8th Parkway 58	devcast.solutions	921418.3831
424	Feedfish	3rd Way 4	feedfish.app	861340.2736
425	Eidel	Forest Avenue 51	eidel.io	562295.6667
426	Wikibox	Ridgeway Circle 73	wikibox.org	657979.3134
427	Flashspan	Drewry Street 30	flashspan.org	911490.4373
428	Realcube	Debra Point 33	realcube.me	994191.952
429	Kwilith	Hintze Parkway 84	kwilith.me	419993.3648
430	Twinte	Sundown Road 76	twinte.support	965487.3799
431	Twinder	Dixon Lane 79	twinder.global	953635.3178
432	Thoughtstorm	Forest Dale Terrace 47	thoughtstorm.site	771806.4431
433	BlogXS	Service Parkway 64	blogxs.online	281652.9807
434	Flashspan	Tony Way 26	flashspan.site	950181.3931
435	Devpoint	Kings Junction 42	devpoint.net	793801.4194
436	Omba	Golf View Junction 18	omba.xyz	650144.6386
437	Meembee	Mockingbird Drive 85	meembee.site	96702.3865
438	InnoZ	Jenna Alley 22	innoz.co	544244.8215
439	Yakidoo	Gina Road 0	yakidoo.online	395220.1235
440	Lazzy	Thackeray Pass 79	lazzy.app	337835.9425
441	Divavu	Beilfuss Trail 85	divavu.org	940527.8338
442	Kimia	Quincy Park 11	kimia.net	25139.3183
443	Jatri	Walton Park 76	jatri.solutions	744942.7226
444	Voomm	Sycamore Way 97	voomm.info	336128.689
445	Meezzy	Northfield Park 41	meezzy.solutions	574490.2223
446	Zoonoodle	Lindbergh Place 77	zoonoodle.global	520957.7939
447	Plambee	Summerview Trail 70	plambee.site	162091.1117
448	Wikido	Arapahoe Street 44	wikido.xyz	333279.19
449	Jayo	Manufacturers Place 36	jayo.io	74360.8983
450	Trudeo	Hagan Street 6	trudeo.online	825607.9008
451	Skynoodle	Hovde Pass 0	skynoodle.global	270784.675
452	Rhycero	Porter Way 2	rhycero.online	372.0011
453	Photofeed	Warbler Circle 97	photofeed.solutions	979243.3842
454	Yambee	Oneill Park 43	yambee.io	684182.9886
455	Meejo	Charing Cross Pass 48	meejo.io	95007.7705
456	Fatz	Shasta Point 48	fatz.solutions	189259.8673
457	Avamba	Lindbergh Parkway 75	avamba.support	116327.7542
458	Kazio	Weeping Birch Road 79	kazio.me	188591.0926
459	Feednation	Graedel Drive 15	feednation.org	778807.9415
460	Dazzlesphere	Hovde Terrace 84	dazzlesphere.site	378350.3906
461	Bluejam	Bartillon Parkway 85	bluejam.site	225785.9631
462	Browsedrive	Di Loreto Junction 60	browsedrive.xyz	545935.3315
463	Skilith	Homewood Junction 59	skilith.app	299400.5046
464	Viva	Carpenter Street 86	viva.me	541885.1185
465	Blognation	Lyons Road 84	blognation.net	49316.8355
466	Skiptube	Corscot Way 72	skiptube.solutions	27439.9843
467	Reallinks	Reindahl Pass 91	reallinks.com	128313.9455
468	Muxo	Hansons Avenue 27	muxo.net	853882.3298
469	Nlounge	Fulton Circle 74	nlounge.org	773682.0053
470	Thoughtsphere	Dawn Pass 54	thoughtsphere.solutions	255000.4872
471	Gigabox	Rusk Drive 85	gigabox.site	634995.8668
472	Flashspan	Jackson Parkway 55	flashspan.solutions	140208.9751
473	Youtags	Londonderry Plaza 76	youtags.me	776640.4969
474	Innotype	Graedel Point 61	innotype.solutions	345077.5981
475	Trudeo	8th Parkway 56	trudeo.co	565004.9323
476	Youspan	Bobwhite Terrace 62	youspan.xyz	964810.5566
477	Skippad	Farragut Way 70	skippad.support	488564.8794
478	Roodel	Dunning Lane 21	roodel.org	819784.892
479	Skivee	Oak Valley Hill 66	skivee.solutions	477635.6749
480	Youfeed	Hooker Alley 25	youfeed.online	174230.8867
481	Flashdog	Grasskamp Crossing 9	flashdog.app	182001.1317
482	JumpXS	Sutteridge Junction 16	jumpxs.global	988622.4795
483	Tazzy	Sunnyside Court 0	tazzy.com	28663.7667
484	Dynazzy	Ridgeway Pass 77	dynazzy.me	206529.921
485	Myworks	Fordem Trail 20	myworks.online	924851.129
486	Skynoodle	Eagle Crest Circle 60	skynoodle.app	459043.6194
487	Blogpad	Upham Park 1	blogpad.com	984405.2708
488	Brainverse	Carberry Park 73	brainverse.co	581125.5587
489	Gabtune	Shasta Way 90	gabtune.io	226100.8405
490	Voolith	American Lane 98	voolith.com	109758.2927
491	Kwinu	Roth Point 95	kwinu.com	431197.8682
492	Skimia	Westend Lane 44	skimia.me	934156.8874
493	Jaloo	Clyde Gallagher Park 83	jaloo.global	173853.0588
494	Eimbee	Arizona Point 85	eimbee.global	55409.2968
495	Avamba	Corben Hill 24	avamba.biz	564821.1523
496	Gevee	Vermont Junction 50	gevee.support	74663.7124
497	Tagopia	Annamark Avenue 65	tagopia.site	504001.9327
498	Talane	Hudson Trail 58	talane.solutions	907330.6542
499	Flashspan	Lukken Circle 42	flashspan.net	347085.7178
500	Browsetype	Lakeland Avenue 8	browsetype.online	143374.8005
\.



COPY wfg.operation (operation_id, name, description) FROM stdin;
1	Deposito	Consiste en agregar dinero a la cuenta
2	Retiro	Consiste en sustraer dinero a la cuenta
3	Transferencia	Consiste en mover dinero de una cuenta a otra
\.

COPY wfg.transaction (transaction_id, account_id, operation_id, mount, date) FROM stdin;
\.



SELECT pg_catalog.setval('wfg.account_account_id_seq', 1, false);



SELECT pg_catalog.setval('wfg.operation_operation_id_seq', 1, false);



SELECT pg_catalog.setval('wfg.transaction_transaction_id_seq', 1, false);



ALTER TABLE ONLY wfg.account
    ADD CONSTRAINT account_pkey PRIMARY KEY (account_id);



ALTER TABLE ONLY wfg.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);



ALTER TABLE ONLY wfg.operation
    ADD CONSTRAINT operation_pkey PRIMARY KEY (operation_id);



ALTER TABLE ONLY wfg.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (transaction_id);



ALTER TABLE ONLY wfg.transaction
    ADD CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES wfg.account(account_id);



ALTER TABLE ONLY wfg.account
    ADD CONSTRAINT fk_account_type FOREIGN KEY (account_type_id) REFERENCES wfg.account_type(account_type_id);



ALTER TABLE ONLY wfg.account
    ADD CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES wfg.customer(customer_id);



ALTER TABLE ONLY wfg.transaction
    ADD CONSTRAINT fk_operation FOREIGN KEY (operation_id) REFERENCES wfg.operation(operation_id);

