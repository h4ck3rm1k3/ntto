package ntto

var DefaultRules = `
dbp         http://dbpedia.org/resource/

dbo         http://dbpedia.org/ontology/
dbprop      http://dbpedia.org/property/

gnd         http://d-nb.info/gnd/
dnb         http://d-nb.info/standards/elementset/gnd#
dnbac       http://d-nb.info/standards/vocab/gnd/geographic-area-code#
dnbvo       http://d-nb.info/standards/vocab/gnd/

viaf        http://viaf.org/viaf/
frbr        http://rdvocab.info/uri/schema/FRBRentitiesRDA/
rdgr        http://rdvocab.info/ElementsGr2/

foaf        http://xmlns.com/foaf/0.1/
rdf         http://www.w3.org/1999/02/22-rdf-syntax-ns#
rdfs        http://www.w3.org/2000/01/rdf-schema#
schema      http://schema.org/
dc          http://purl.org/dc/elements/1.1/
dcterms     http://purl.org/dc/terms/


# generic freebase
fb          http://rdf.freebase.com/ns/
fbkey       http://rdf.freebase.com/key/

rdfa        http://www.w3.org/ns/rdfa#
virtrdf     http://www.openlinksw.com/virtrdf-data-formats#
umbel       http://umbel.org/umbel#
umbelac     http://umbel.org/umbel/ac/
umbelsc     http://umbel.org/umbel/sc/
prov        http://www.w3.org/ns/prov#

# wikidata
wd          http://www.wikidata.org/entity/
wdo         http://www.wikidata.org/ontology#

# more dbpedia languages w/ > 100k pages
# dbpedia languages (more below)
dbpde   http://de.dbpedia.org/resource/
dbpfr   http://fr.dbpedia.org/resource/
dbpen   http://en.dbpedia.org/resource/
dbpes   http://es.dbpedia.org/resource/
dbpit   http://it.dbpedia.org/resource/
dbpnl   http://nl.dbpedia.org/resource/
dbpru   http://ru.dbpedia.org/resource/
dbpsv   http://sv.dbpedia.org/resource/
dbppl   http://pl.dbpedia.org/resource/
dbpja   http://ja.dbpedia.org/resource/
dbppt   http://pt.dbpedia.org/resource/
dbpar   http://ar.dbpedia.org/resource/
dbpzh   http://zh.dbpedia.org/resource/
dbpuk   http://uk.dbpedia.org/resource/
dbpca   http://ca.dbpedia.org/resource/
dbpno   http://no.dbpedia.org/resource/
dbpfi   http://fi.dbpedia.org/resource/
dbpcs   http://cs.dbpedia.org/resource/
dbphu   http://hu.dbpedia.org/resource/
dbptr   http://tr.dbpedia.org/resource/
dbpro   http://ro.dbpedia.org/resource/
dbpsw   http://sw.dbpedia.org/resource/
dbpko   http://ko.dbpedia.org/resource/
dbpkk   http://kk.dbpedia.org/resource/
dbpvi   http://vi.dbpedia.org/resource/
dbpda   http://da.dbpedia.org/resource/
dbpeo   http://eo.dbpedia.org/resource/
dbpsr   http://sr.dbpedia.org/resource/
dbpid   http://id.dbpedia.org/resource/
dbplt   http://lt.dbpedia.org/resource/
dbpvo   http://vo.dbpedia.org/resource/
dbpsk   http://sk.dbpedia.org/resource/
dbphe   http://he.dbpedia.org/resource/
dbpfa   http://fa.dbpedia.org/resource/
dbpbg   http://bg.dbpedia.org/resource/
dbpsl   http://sl.dbpedia.org/resource/
dbpeu   http://eu.dbpedia.org/resource/
dbpwar  http://war.dbpedia.org/resource/
dbpet   http://et.dbpedia.org/resource/
dbphr   http://hr.dbpedia.org/resource/
dbpms   http://ms.dbpedia.org/resource/
dbphi   http://hi.dbpedia.org/resource/
dbpsh   http://sh.dbpedia.org/resource/

dbpropde   http://de.dbpedia.org/property/
dbpropfr   http://fr.dbpedia.org/property/
dbpropen   http://en.dbpedia.org/property/
dbpropes   http://es.dbpedia.org/property/
dbpropit   http://it.dbpedia.org/property/
dbpropnl   http://nl.dbpedia.org/property/
dbpropru   http://ru.dbpedia.org/property/
dbpropsv   http://sv.dbpedia.org/property/
dbproppl   http://pl.dbpedia.org/property/
dbpropja   http://ja.dbpedia.org/property/
dbproppt   http://pt.dbpedia.org/property/
dbpropar   http://ar.dbpedia.org/property/
dbpropzh   http://zh.dbpedia.org/property/
dbpropuk   http://uk.dbpedia.org/property/
dbpropca   http://ca.dbpedia.org/property/
dbpropno   http://no.dbpedia.org/property/
dbpropfi   http://fi.dbpedia.org/property/
dbpropcs   http://cs.dbpedia.org/property/
dbprophu   http://hu.dbpedia.org/property/
dbproptr   http://tr.dbpedia.org/property/
dbpropro   http://ro.dbpedia.org/property/
dbpropsw   http://sw.dbpedia.org/property/
dbpropko   http://ko.dbpedia.org/property/
dbpropkk   http://kk.dbpedia.org/property/
dbpropvi   http://vi.dbpedia.org/property/
dbpropda   http://da.dbpedia.org/property/
dbpropeo   http://eo.dbpedia.org/property/
dbpropsr   http://sr.dbpedia.org/property/
dbpropid   http://id.dbpedia.org/property/
dbproplt   http://lt.dbpedia.org/property/
dbpropvo   http://vo.dbpedia.org/property/
dbpropsk   http://sk.dbpedia.org/property/
dbprophe   http://he.dbpedia.org/property/
dbpropfa   http://fa.dbpedia.org/property/
dbpropbg   http://bg.dbpedia.org/property/
dbpropsl   http://sl.dbpedia.org/property/
dbpropeu   http://eu.dbpedia.org/property/
dbpropwar   http://war.dbpedia.org/property/
dbpropet   http://et.dbpedia.org/property/
dbprophr   http://hr.dbpedia.org/property/
dbpropms   http://ms.dbpedia.org/property/
dbprophi   http://hi.dbpedia.org/property/
dbpropsh   http://sh.dbpedia.org/property/

address     http://schemas.talis.com/2005/address/schema#
admin       http://webns.net/mvcb/
atom        http://atomowl.org/ontologies/atomrdf#
atom        http://www.w3.org/2005/Atom
aws         http://soap.amazon.com/
b3s         http://b3s.openlinksw.com/
batch       http://schemas.google.com/gdata/batch/
bibo        http://purl.org/ontology/bibo/
bugzilla    http://www.openlinksw.com/schemas/bugzilla#
c           http://www.w3.org/2002/12/cal/icaltzd#
category    http://dbpedia.org/resource/Category:
cb          http://www.crunchbase.com/
cc          http://web.resource.org/cc/
content     http://purl.org/rss/1.0/modules/content/
cv          http://purl.org/captsolo/resume-rdf/0.2/cv#
cvbase      http://purl.org/captsolo/resume-rdf/0.2/base#
dawgt       http://www.w3.org/2001/sw/DataAccess/tests/test-dawg#
digg        http://digg.com/docs/diggrss/
ebay        urn:ebay:apis:eBLBaseComponents
enc         http://purl.oclc.org/net/rss_2.0/enc#
exif        http://www.w3.org/2003/12/exif/ns/
facebook    http://api.facebook.com/1.0/
ff          http://api.friendfeed.com/2008/03
fn          http://www.w3.org/2005/xpath-functions/#
g           http://base.google.com/ns/1.0/
gb          http://www.openlinksw.com/schemas/google-base#
gd          http://schemas.google.com/g/2005/
geo         http://www.w3.org/2003/01/geo/wgs84_pos#
geonames    http://www.geonames.org/ontology#
georss      http://www.georss.org/georss/
gml         http://www.opengis.net/gml/
go          http://purl.org/obo/owl/GO#
grs         http://www.georss.org/georss/
hlisting    http://www.openlinksw.com/schemas/hlisting/
hoovers     http://wwww.hoovers.com/
hrev        http:/www.purl.org/stuff/hrev#
ical        http://www.w3.org/2002/12/cal/ical#
ir          http://web-semantics.org/ns/image-regions/
itunes      http://www.itunes.com/DTDs/Podcast-1.0.dtd
lgv         http://linkedgeodata.org/vocabulary#
link        http://www.xbrl.org/2003/linkbase/
lod         http://lod.openlinksw.com/
math        http://www.w3.org/2000/10/swap/math#
media       http://search.yahoo.com/mrss/
mesh        http://purl.org/commons/record/mesh/
meta        urn:oasis:names:tc:opendocument:xmlns:meta:1.0
mf          http://www.w3.org/2001/sw/DataAccess/tests/test-manifest#
mmd         http://musicbrainz.org/ns/mmd-1.0#
mo          http://purl.org/ontology/mo/
mql         http://www.freebase.com/
nci         http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#
nfo         http://www.semanticdesktop.org/ontologies/nfo/#
ng          http://www.openlinksw.com/schemas/ning#
nyt         http://www.nytimes.com/
oai         http://www.openarchives.org/OAI/2.0/
oai_dc      http://www.openarchives.org/OAI/2.0/oai_dc/
obo         http://www.geneontology.org/formats/oboInOwl#
office      urn:oasis:names:tc:opendocument:xmlns:office:1.0
oo          urn:oasis:names:tc:opendocument:xmlns:meta:1.0:
openSearch  http://a9.com/-/spec/opensearchrss/1.0/
opl         http://www.openlinksw.com/schema/attribution#
opl-gs      http://www.openlinksw.com/schemas/getsatisfaction/
opl-meetup  http://www.openlinksw.com/schemas/meetup/
opl-xbrl    http://www.openlinksw.com/schemas/xbrl/
oplweb      http://www.openlinksw.com/schemas/oplweb#
ore         http://www.openarchives.org/ore/terms/
owl         http://www.w3.org/2002/07/owl#
product     http://www.buy.com/rss/module/productV2/
protseq     http://purl.org/science/protein/bysequence/
r           http://backend.userland.com/rss2/
radio       http://www.radiopop.co.uk/
rev         http://purl.org/stuff/rev#
review      http:/www.purl.org/stuff/rev#
rss         http://purl.org/rss/1.0/
sc          http://purl.org/science/owl/sciencecommons/
scovo       http://purl.org/NET/scovo#
sf          urn:sobject.enterprise.soap.sforce.com
sioc        http://rdfs.org/sioc/ns#
sioct       http://rdfs.org/sioc/types#
skos        http://www.w3.org/2004/02/skos/core#
slash       http://purl.org/rss/1.0/modules/slash/
stock       http://xbrlontology.com/ontology/finance/stock_market#
twfy        http://www.openlinksw.com/schemas/twfy#
uniprot     http://purl.uniprot.org/
usc         http://www.rdfabout.com/rdf/schema/uscensus/details/100pct/
v           http://www.openlinksw.com/xsltext/
vcard       http://www.w3.org/2001/vcard-rdf/3.0#
vcard2006   http://www.w3.org/2006/vcard/ns#
vi          http://www.openlinksw.com/virtuoso/xslt/
virt        http://www.openlinksw.com/virtuoso/xslt/
virtcxml    http://www.openlinksw.com/schemas/virtcxml#
virtrdf     http://www.openlinksw.com/schemas/virtrdf#
void        http://rdfs.org/ns/void#
wb          http://www.worldbank.org/
wf          http://www.w3.org/2005/01/wf/flow#
wfw         http://wellformedweb.org/CommentAPI/
xf          http://www.w3.org/2004/07/xpath-functions/
xfn         http://gmpg.org/xfn/11#
xhtml       http://www.w3.org/1999/xhtml/
xhv         http://www.w3.org/1999/xhtml/vocab#
xi          http://www.xbrl.org/2003/instance/
xml         http://www.w3.org/XML/1998/namespace/
xn          http://www.ning.com/atom/1.0/
xsd         http://www.w3.org/2001/XMLSchema#
xsl10       http://www.w3.org/XSL/Transform/1.0
xsl1999     http://www.w3.org/1999/XSL/Transform/
xslwd       http://www.w3.org/TR/WD-xsl/
y           urn:yahoo:maps
yago        http://dbpedia.org/class/yago/
yt          http://gdata.youtube.com/schemas/2007/
zem         http://s.zemanta.com/ns#
`
