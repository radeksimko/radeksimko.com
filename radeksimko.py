import webapp2

class MainPage( webapp2.RequestHandler ):
	def get( self ):
		self.redirect( uri='http://about.radeksimko.com' )

class LinkedIn( webapp2.RequestHandler ):
	def get( self ):
		self.redirect( uri='http://linkedin.com/in/radeksimko' )

class PlainText( webapp2.RequestHandler ):
	def get( self ):
		self.response.headers['Content-Type'] = 'text/plain'

app = webapp2.WSGIApplication([
	('/robots.txt', PlainText),
	('/pinterest-24f5e.html', PlainText),
	('/attachments/cv-radek-simko.pdf', LinkedIn),
	('/.*', MainPage)
], debug=True)