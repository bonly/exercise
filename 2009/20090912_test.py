#------------------------------------------------------------------------------
#           Name: test.py
#         Author: Kevin Harris
#  Last Modified: 04/29/05
#    Description: 
#------------------------------------------------------------------------------

import extendAndEmbedTest

#------------------------------------------------------------------------------
# Define a function and set it as a call-back for the application to use
#------------------------------------------------------------------------------

def somePythonCallBackFunction( n ):
    print "Application called \"somePythonCallBackFunction\" and passed: " + str( n )

extendAndEmbedTest.setCallback( somePythonCallBackFunction );

#------------------------------------------------------------------------------
# Test the getValue() function which gets an int value form the application. 
#------------------------------------------------------------------------------

print "extendAndEmbedTest.getApplicationValue() returned: " + str( extendAndEmbedTest.getApplicationValue() )

#------------------------------------------------------------------------------
# Test the setApplicationValue() function which sets an int value in the 
# application.
#------------------------------------------------------------------------------

extendAndEmbedTest.setApplicationValue( 55 )

# Let's check it again...
print "extendAndEmbedTest.getApplicationValue() returned: " + str( extendAndEmbedTest.getApplicationValue() )
	
#------------------------------------------------------------------------------
# Set a return value so we can extract it after the script has executed.
#------------------------------------------------------------------------------

returnValue = -1

